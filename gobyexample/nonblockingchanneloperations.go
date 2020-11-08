// basic sends and receives on channels are blocking. However, select can be used with a DEFAUL clause to implement non-blocking sends, receivesm multiway selects

package main

import"fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// this is a non-blocking receive. If a value is available on messages then select will take the <- messages case with that value. Or else it will take the default case

	select {
	case msg := <- messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// a non-blocking send works similarly. Here msg cannot be sent to the messages channel becasue the channel has no buffer and there is no receiver. Therefore the DEFAULT case is selected
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

// we can use multiple cases above the DEFAULT clause to implement a multiway nonblocking select. Here we attempt nonblocknig receives on both messages and signals
	select {
	case msg:= <- messages:
		fmt.Println("received message", msg)
	case sig := <- signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}