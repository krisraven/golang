// timeouts are important for programs that connect to external resources or that need to bound execution time. Select and Channels make implementing timeouts in Go easy and elegant

package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

// this select is implementing the timeout. res:=<-c1 awaits the result and <-time.After awaits a value to be sent after 1 second (but it actually takes 2 seconds to complete)
	select {
	case res := <- c1:
		fmt.Println(res)
	case <- time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <- c2:
		fmt.Println(res)
	case <- time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
// running this program it shows the first operation timing out and the second succeeding