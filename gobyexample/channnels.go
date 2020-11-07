package main 

import "fmt"
// channels are the pipes thtat connect concurrent goroutines. You can send values into chanels from one goroutine and receive those values into another goroutine

func main() {
	messages := make(chan string) // this creates a new channel. Channels are typed by the values they convey

	go func() { messages <- "ping" }() // this sends a value "ping" to the messages channel (using <- syntax)

	msg := <- messages
	fmt.Println(msg)
} 

// by default send and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronisation