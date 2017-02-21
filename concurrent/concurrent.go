// Package concurrent is a toy demo of Go standard concurrency
// Based on 2012 talk 'Go Concurrency Patterns' by Rob Pike
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {

	c := make(chan string)

	go boring("message", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("Channel val: %q\n", <-c)
	}
	fmt.Println("Exiting...")
}
