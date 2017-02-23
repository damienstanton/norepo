package intf

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type ClientStruct struct {
	Token     string
	Backends  []string
	Backoff   time.Duration
	Tolerance int
	robin     uint64
}

// consider something like
func (c *ClientStruct) Do(r *http.Request) (res *http.Response, err error) {
	r.Header.Add("Authorization", c.Token)
	for i := 0; i < c.Tolerance; i++ {
		r.URL.Host = c.Backends[atomic.AddUint64(&c.robin, 1)%uint64(len(c.Backends))]
		log.Printf("%s: %s %s", r.UserAgent(), r.Method, r.URL)
		res, err = http.DefaultClient.Do(r)
		if err != nil {
			time.Sleep(time.Duration(i) * c.Backoff)
			continue
		}
		break
	}
	return res, err
}

// ClientInterface shows implementation as a single-method interface.
// Single-method interfaces are to be sought, by reason to follow below.
type ClientInterface interface {
	Do(*http.Request) (*http.Response, error)
}

// ClientFunc implements ClientInterface
type ClientFunc func(*http.Request) (*http.Response, error)

// Because any type can have methods, function types
// can implement single-method interfaces with a self-call.
func (f ClientFunc) Do(r *http.Request) (*http.Response, error) { return f(r) }

// Decorator pattern is dead simple
type Decorator func(ClientInterface) ClientInterface

// and this can be useful, if dense
func Logging(l *log.Logger) Decorator {
	return func(c ClientInterface) ClientInterface {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			l.Printf("%s: %s %s", r.UserAgent(), r.Method, r.URL)
			return c.Do(r)
		})
	}
}

// dummy instrumentation client structures
type (
	Counter   struct{}
	Histogram struct{}
)

// here for instrumentation instead of logging
func Instrumentation(requests Counter, latency Histogram) Decorator {
	return func(c ClientInterface) ClientInterface {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			defer func(start time.Time) {
			}(time.Now())
			return c.Do(r)
		})
	}
}

// FaultTolerance, and so on
func FaultTolerance(attempts int, backoff time.Duration) Decorator {
	return func(c ClientInterface) ClientInterface {
		return ClientFunc(func(r *http.Request) (res *http.Response, err error) {
			for i := 0; i <= attempts; i++ {
				if res, err = c.Do(r); err == nil {
					break
				}
				time.Sleep(backoff * time.Duration(i))
			}
			return res, err
		})
	}
}

// can extend to other areas, like auth
func Authorization(token string) Decorator {
	return Header("Authorization", token)
}

func Header(name, value string) Decorator {
	return func(c ClientInterface) ClientInterface {
		return ClientFunc(func(r *http.Request) (*http.Response, error) {
			r.Header.Add(name, value)
			return c.Do(r)
		})
	}
}

// finally, for general composition
func Decorate(c ClientInterface, ds ...Decorator) ClientInterface {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
} // see main.go for use
