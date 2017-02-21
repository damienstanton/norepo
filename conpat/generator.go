package main

import (
	"fmt"
	"time"
)

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
	// select makes this better
	go func() {
		for {
			select {
			case s := <-in1:
				c <- s
			case s := <-in2:
				c <- s
			}
		}
	}()
	return c
}

// whisper represents the 'Chinese Whispers Game' section of the talk
func whisper(left, right chan int) {
	// neato
	left <- 1 + <-right
}
