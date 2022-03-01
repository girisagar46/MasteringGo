package main

import (
	"fmt"
)

/*
Basic sends and receives on channels are blocking.
However, we can use select with a default clause to implement non-blocking sends, receives, and even non-blocking multi-way selects.
*/

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages: // there's no value in the messages channel, so this case fails
		fmt.Println("received message", msg)
	default:
		fmt.Println("No message received")
	}

	msg := "hi"
	select {
	case messages <- msg: //  Here msg cannot be sent to the messages channel, because the channel has no buffer and there is no receiver.
		fmt.Println("sent message")
	default:
		fmt.Println("no message sent")
	}

	select { // select with multiple cases before default
	case msg := <-messages: // there's nothing in the messages channel
		fmt.Println("received message", msg)
	case sig := <-signals: // there's nothing in the signals channel
		fmt.Println("received signal", sig)
	default: // hence, default case is selected
		fmt.Println("No activity")
	}
}

/*
No message received
no message sent
No activity
*/
