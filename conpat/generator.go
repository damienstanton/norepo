package main

import (
	"fmt"
	"time"
)

/* Generator: function that returns a channel */

// boringGenerator returns a receive-only chan of strings.
// This one is nice because simply different invocations can describe distinct
// services. For example: boringGenerator("A") boringGenerator("B")
func boringGenerator(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(SleepVal)
		}
	}()
	// now return the chan to the caller
	return c
}

// fanIn is a demo multiplexer
func fanIn(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	// these could easily be inlined for tidiness, but this is good to see.
	go func() {
		for {
			c <- <-in1
		}
	}()
	go func() {
		for {
			c <- <-in2
		}
	}()
	return c
}
