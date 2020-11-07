// channels can be used to synchronise execution across goroutines
package main

import (
	"fmt"
	"time"
)
// this function will run in a goroutine. The done channel will be used to notify another goroutine that this functions work is done
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
// start a woker and give it a channel to notify on
	done := make(chan bool, 1)
	go worker(done)
// block untuk we receive a notification from the worker on the channel
	<- done
}