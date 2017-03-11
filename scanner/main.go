package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t string

	scanner := bufio.NewScanner(os.Stdin)
	for t != ":q" {
		fmt.Print(">> ")
		scanner.Scan()
		t = scanner.Text()
		if t != ":q" {
			fmt.Printf("OK. Received: %v\n", t)
		}
		if t == ":q" {
			fmt.Println("OK. Quit signal received. Bye!")
		}
	}
}
