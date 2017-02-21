// Program conpat is based on 2012 talk 'Go Concurrency Patterns' by Rob Pike
// as well as the 2013 talk 'Advanced Go Concurrency Patterns' by Sameer Ajmani
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SleepVal is how long we want our fake exec time to be
var SleepVal = time.Duration(rand.Intn(1e3)) * time.Millisecond

// Ball is the ball in a ping-pong game
type Ball struct {
	hits int
}

func main() {
	/* Basics */

	fmt.Println("Running fanIn example...")
	c := fanIn(boringGenerator("A"), boringGenerator("B"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Exiting both fanIn goroutines...")
	fmt.Println("Running daisy chain example...")
	// goroutines are lightweight, bro
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
	// use the First func, and time it
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := First("golang", fakeSearch("rep 1"), fakeSearch("rep 2"))
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)

	/* Advanced */

	// Ping pong game
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)
	// start the game
	table <- new(Ball)
	time.Sleep(1 * time.Second)
	// end the game
	<-table
}
