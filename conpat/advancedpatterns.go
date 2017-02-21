package main

import (
	"fmt"
	"time"
)

// player is one of the players in our ping-pong ball game
func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
