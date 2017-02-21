// Program boltb might show some Bolt functionality, not sure yet
package main

import (
	"fmt"

	"github.com/boltdb/bolt"
)

var dummyObject = []byte("blah")

func main() {
	// Bolt is a COW B+ tree in the file optimized for reads
	// LSM tree in, say LevelDB is optimized for writes
	db, err := bolt.Open("data.db", 0644, nil)
	if err != nil {
		fmt.Errorf("could not make new bolt database: %v", err)
	}
	defer db.Close()

	key := []byte("A")
	val := []byte("1")

	// store data (throwing out errors dangerously)
	_ = db.Update(func(tx *bolt.Tx) error {
		bucket, _ := tx.CreateBucketIfNotExists(dummyObject)
		err = bucket.Put(key, val)
		if err != nil {
			return err
		}
		return nil
	})

	// get data (throwing out errors dangerously)
	_ = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(dummyObject)
		if bucket == nil {
			return fmt.Errorf("No bucket %q", dummyObject)
		}
		val := bucket.Get(key)
		fmt.Println(string(val))
		return nil
	})
}
