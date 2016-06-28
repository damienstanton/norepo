(ns progenitor.core)

;; This is blatantly copied from https://gist.github.com/sindux/e6ec87a88da99c9576a5
;; protobuf class has "hasXXX"/"getXXX" and "getXXXCount"/"getXXXList"
(defn get-methods [^Class class]
  (let [methods (into {} (map (fn [^java.lang.reflect.Method m] [(.getName m) m]) (.getMethods class)))
        has-XXX (into {} (for [[^String name ^java.lang.reflect.Method method] methods
                               :when (= class (.getDeclaringClass method))
                               :let [[field-name has-val-test ^java.lang.reflect.Method getter :as fg]
                                     (cond
                                       (.startsWith name "has")
                                       (let [n (.substring name 3)]
                                         [n identity (methods (str "get" n))])

                                       (and (.startsWith name "get")
                                            (.endsWith name "Count"))
                                       (let [n (.substring name 3 (- (count name) 5))]
                                         [n pos? (methods (str "get" n "List"))])

                                       :else [nil nil nil])]
                               :when (every? (comp not nil?) fg)]
                           [field-name [(.getName method) has-val-test (.getName getter)]]))]
    has-XXX))

(defprotocol ToMap
  (to-map [this] "Convert given object to edn"))

(defn coerce-val [val]
  (cond
    (instance? java.lang.Enum val) (keyword (str
                                             (.getSimpleName (.getClass ^java.lang.Enum val))
                                             "/"
                                             (.toString ^java.lang.Enum val)))
    (instance? java.lang.Boolean val) (boolean val)
    (instance? java.util.List val) (mapv coerce-val val)
    (satisfies? ToMap val) (to-map val)
    :else val))

(defmacro gen-has-get [map class this [field-name [has-m has-val-test get-m]]]
  (let [class (resolve class)]
    `(if (~has-val-test (~(symbol (str "." has-m)) ~(with-meta this {:tag class})))
       (assoc! ~map ~(keyword field-name)
               (coerce-val (~(symbol (str "." get-m)) ~(with-meta this {:tag class}))))
       ~map)))


(defmacro gen-has-gets [map class this methods]
  `(let ~(vec (interleave (repeat map)
                          (for [method# methods]
                            `(gen-has-get ~map ~class ~this ~method#))))
     ~map))

(defmacro gen-to-map [class]
  (let [methods (get-methods (resolve class))]
    `(extend-type ~class
       ToMap
       (to-map [this#]
         (let [res# (transient {})]
           (persistent! (gen-has-gets res# ~class this# ~methods)))))))

(defmacro doseq-m
  [coll body]
  (let [fns (map #(list body (symbol (.getName ^Class %))) (eval coll))]
    `(do ~@fns)))




(defmacro doseq-all
  "Like doseq but eval the expression at compile time (only 1 expr-binding & 1 body is supported) "
  [seq-exprs & body]
  (let [[sym coll] (take 2 seq-exprs)   ;; only support 1 expr-binding
        body (first body)               ;; only support 1 body
        outputs (->> (eval coll)
                     (map #(subst body sym %)))]
    `(do ~@outputs)))
(defn- enumerate-classes
  [package]
  (let [classes (-> (.. Thread currentThread getContextClassLoader)
                     com.google.common.reflect.ClassPath/from
                     (.getTopLevelClasses (str package)))]
    (map #(-> ^com.google.common.reflect.ClassPath$ClassInfo % .getName symbol) classes)))

(defmacro for-all-classes
  "For each classes found in given package, emit given (body)"
  [body package]
  (let [fns (->> (enumerate-classes package)
                 (map #(list body %)))]
    `(do ~@fns)))


;;;;;;;;;; scaffolding ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
(defn suggest-val [^Class type]
  (if (isa? type Enum)
    (map str (java.util.EnumSet/allOf type))
    (condp = type
      Long/TYPE 0
      Integer/TYPE 0
      Double/TYPE 0.0
      Boolean/TYPE false
      String ""
      java.util.Date '(java.util.Date.)
      type)))

(defn scaffold-msg [^Class msg]
  (into (sorted-map)
        (for [^java.lang.reflect.Method m (.getMethods msg)
              :let [method-name (.getName m)]
              :when (and (.startsWith method-name "get")
                         (not (.endsWith method-name "AsLong"))
                         (= 0 (.getParameterCount m))
                         (not (contains? #{"getExcludeMe"} method-name)))
              :let [type (.getReturnType m)
                    type (if (.contains (.getName type) ".groups.")
                           (scaffold-msg type)
                           type)
                    type-val (suggest-val type)]]
          [(keyword (.substring method-name 3)) type-val])))



;;;;;;;;;;;;;; outbound converter (edn->object) ;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

(defn identity-coerce-helper [field-name type obj]
  obj)

(defn try-coerce
  ([type obj] (try-coerce type obj nil identity-coerce-helper))
  ([type obj field-name coerce-helper]
   (cond
     (isa? type Enum) (try
                        {:type type :val (Enum/valueOf type obj)}
                        (catch Exception ex
                          {:type nil :val obj :err-params [type obj] :ex ex}))
     (isa? type java.util.EnumSet) (coerce-helper field-name type obj)
     (and (= #'long type)
          (number? obj)) {:type Long/TYPE :val obj}
     (and (= #'double type)
          (number? obj)) {:type Double/TYPE :val obj}
     (and (= #'boolean type)) {:type Boolean/TYPE :val (boolean obj)}
     (and (class? type)
          (instance? type obj)) {:type type :val obj}
     :else {:type nil :val obj})))

(defn try-coerce-params
  ([param-types params] (try-coerce-params param-types params nil identity-coerce-helper))
  ([param-types params field-name coerce-helper]
   (let [coerced (for [[param-type param] (map vector param-types params)
                       :let [param-type (resolve param-type)
                             coerced (try-coerce param-type param field-name coerce-helper)]]
                   coerced)]
     (when (every? (comp not nil? :type) coerced)
       coerced))))

(defn get-candidate-method
  ([obj method params] (get-candidate-method obj method params identity-coerce-helper))
  ([obj method params coerce-helper]
   (let [candidate-methods (for [member (:members (clojure.reflect/reflect obj))
                                 :when  (and (= (:name member) (symbol method))
                                             (= (count params) (count (:parameter-types member)))
                                             (contains? (:flags member) :public)
                                             (not (contains? (:flags member) :synthetic)))
                                 :let [coerced (try-coerce-params
                                                (:parameter-types member) params
                                                (:name member)  coerce-helper)]
                                 :when (not (nil? coerced))
                                 ]
                             {:method member
                              :coerced-params coerced})]
     candidate-methods
     (if (= 1 (count candidate-methods))
       (first candidate-methods)
       (throw (IllegalArgumentException. (format "Cannot find method %s%s on %s, candidate %s"
                                                 method params obj (pr-str candidate-methods))))))))


(defn invoke
  "coerce-helper is (fn [field-name param-type param-val]) to help coerce param-val to param-type"
  ([^Object obj method] (invoke obj method nil identity-coerce-helper))
  ([^Object obj method params] (invoke obj method params identity-coerce-helper))
  ([^Object obj method params coerce-helper]
   (let [c (.getClass obj)
         coerced (->> (get-candidate-method obj method params coerce-helper) :coerced-params)
         ;;m (.getMethod c method (into-array Class (map :type coerced)))
         p (into-array Object (map :val coerced))]
     #_(.invoke m obj (to-array params))
     (clojure.lang.Reflector/invokeInstanceMethod obj method p))))
