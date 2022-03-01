package main

import (
	"fmt"
	"time"
)

/*
Go’s select lets you wait on multiple channel operations.
Combining goroutines and channels with select is a powerful feature of Go.
*/

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second) // simulate e.g. blocking RPC operations executing in concurrent goroutines.
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second) // simulate e.g. blocking RPC operations executing in concurrent goroutines.
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select { // use select to await both of these values simultaneously, printing each one as it arrives.
		case msg1 := <-c1:
			fmt.Println("Received from c1", msg1)
		case msg2 := <-c2:
			fmt.Println("Received from c2", msg2)
		}
	}
}

// NOTE: total execution time is only ~2 seconds since both the 1 and 2 second Sleeps execute concurrently.

/*
$ time go run 29select.go
Received from c1 one
Received from c2 two
go run 29select.go  0.13s user 0.17s system 11% cpu 2.552 total
*/
