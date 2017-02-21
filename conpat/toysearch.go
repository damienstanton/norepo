package main

import (
	"fmt"
	"time"
)

var (
	Web1   = fakeSearch("web")
	Image1 = fakeSearch("Image")
	Video1 = fakeSearch("Video")
	Web2   = fakeSearch("web")
	Image2 = fakeSearch("Image")
	Video2 = fakeSearch("Video")
)

// Result is a fake result
type Result string

// Search is our simulated search data structure
type Search func(query string) Result

func fakeSearch(k string) Search {
	return func(q string) Result {
		time.Sleep(SleepVal)
		return Result(fmt.Sprintf("%s result for %q\n", k, q))
	}
}

// Googlev10 does a web search on that site
func GoogleV10(q string) (res []Result) {
	res = append(res, Web(q))
	res = append(res, Image(q))
	res = append(res, Video(q))
	return res
}

// Googlev20 does a better web search on that site
func GoogleV20(q string) (finalResults []Result) {
	c := make(chan Result)
	// concurrently rather than serialized
	go func() {
		c <- Web(q)
	}()
	go func() {
		c <- Image(q)
	}()
	go func() {
		c <- Video(q)
	}()

	for i := 0; i < 3; i++ {
		res := <-c
		finalResults = append(finalResults, res)
	}
	return finalResults
}

// Googlev21 does a betterweb search on that site
func GoogleV21(q string) (finalResults []Result) {
	c := make(chan Result)
	go func() {
		c <- Web(q)
	}()
	go func() {
		c <- Image(q)
	}()
	go func() {
		c <- Video(q)
	}()

	// now no need to wait for slow servers. Still no locks, condition vars,
	// or callbacks. This is nice.
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case res := <-c:
			finalResults = append(finalResults, res)
		case <-timeout:
			fmt.Println("Timed out")
			return
		}
	}
	return finalResults
}

// First replicates the servers
// Goal is to avoid throwing out the results if a timeout expires
func First(q string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](q) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

// GoogleV30 sums the feature set we've been iterating on
// It's good because we can now guarantee that all 3 results return within 80ms
// from a set of replica servers.
func GoogleV30(q string) (finalResults []Result) {
	c := make(chan Result)
	go func() {
		c <- First(q, Web1, Web2)
	}()
	go func() {
		c <- First(q, Image1, Image2)
	}()
	go func() {
		c <- First(q, Video1, Video2)
	}()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case res := <-c:
			finalResults = append(finalResults, res)
		case <-timeout:
			fmt.Println("Timed out")
			return
		}
	}
	return finalResults
}
