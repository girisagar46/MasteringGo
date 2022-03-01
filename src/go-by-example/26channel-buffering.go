package main

import "fmt"

/*
Buffered channels accept a limited number of values without a corresponding receiver for those values.
*/
func main() {
	// make(chan string) will make unbuffered channel
	messages := make(chan string, 2) // channel of strings buffering upto 2 value

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	messages <- "Hello"
	messages <- "World"
	//messages <- "Deadlock" // fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	//fmt.Println(<-messages) // fatal error: all goroutines are asleep - deadlock!

}
