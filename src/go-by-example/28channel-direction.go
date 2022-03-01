package main

import "fmt"

// This ping function only accepts a channel for sending values.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings // receive the msg from pings channel
	pongs <- msg   // send the msg to pongs channel
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings /*sender channel*/, "message")
	pong(pings /*sender channel*/, pongs /*receiver channel*/)
	fmt.Println("message in pongs channel = ", <-pongs)
}

/*
message in pongs channel =  message
*/
