package main

import (
	"fmt"
)

/*
Channels are the pipes that connect concurrent goroutines.
You can send values into channels from one goroutine and receive those values into another goroutine.

By default, sends and receives block until both the sender and receiver are ready.
*/

func main() {
	messages := make(chan string) // Create a new channel with make(chan val-type). Channels are typed by the values they convey.

	// Send a value into a channel using the `channel <-` syntax.
	go func() { messages <- "Ping" }() // sending a value to channel via new goroutine

	// The `<-channel` syntax receives a value from the channel.
	msg := <-messages
	fmt.Println(msg)
}
