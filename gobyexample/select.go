// gos select let you wait on multiple channel operations. Combing goroutines and channels with select is a powerful feature of go

package main

import (
	"fmt"
	"time"
)

func main() {
// this example will select across 2 channels
	c1 := make(chan string)
	c2 := make(chan string)

// each channel will receive a value after some amount of time to simulate blocking RPC operations executing in concurrent goroutines
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

// select is used to await both of these values, printing each one as it arrives
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- c1:
				fmt.Println("received", msg1)
		case msg2 := <- c2:
			fmt.Println("received", msg2)
		}
	}
}