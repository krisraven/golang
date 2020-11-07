// by default channels are unbuffered meaning the will only accept sends (chan <-) if there is a corresponding receive ready to receive the sent value. Buffered channels accept a limited number of values wit ha corresponding receiver for those values 

package main

import "fmt"

func main() {

	messages := make(chan string, 2) // here we make a channel of up to 2 strings

	// because this channel is buffered we can send these values into the channel without a corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<- messages)
	fmt.Println(<- messages)
}