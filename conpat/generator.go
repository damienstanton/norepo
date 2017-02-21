package main

import (
	"fmt"
	"time"
)

/* Generator: function that returns a channel */

// boringGenerator returns a receive-only chan of strings
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
