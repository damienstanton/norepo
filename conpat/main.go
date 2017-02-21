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
	fmt.Println("Running fanIn example...")
	c := fanIn(boringGenerator("A"), boringGenerator("B"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Exiting both fanIn goroutines...")
	fmt.Println("Running daisy chain example...")
	const n = 100000
	// min in this case is the "leftmost" agent
	min := make(chan int)
	// initialize our starting position
	right := min
	left := min
	for i := 0; i < n; i++ {
		right = make(chan int)
		go whisper(left, right)
		left = right
	}
	// move down the chain with an anonymous func
	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-min)
}
