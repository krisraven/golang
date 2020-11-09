// for and range provide iteration over basic data structures. We can also use this syntax to iterate over values received from a channel

package main

import "fmt"

func main() {
	// we can iterate over 2 values in the queue channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// this range terates over each element as its received from queue. Because we closed the channel above the iteration terminates after receiving the 2 elements
	for elem := range queue {
		fmt.Println(elem)
	}
}

// this example shows that it's possible toclose a non-empty channel but still have the remaining values be received
