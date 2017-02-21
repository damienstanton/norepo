// Package conpat is based on 2012 talk 'Go Concurrency Patterns' by Rob Pike
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SleepVal is how long we want our fake exec time to be
var SleepVal = time.Duration(rand.Intn(1e3)) * time.Millisecond

func main() {
	c := fanIn(boringGenerator("A"), boringGenerator("B"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Exiting both goroutines...")
}
