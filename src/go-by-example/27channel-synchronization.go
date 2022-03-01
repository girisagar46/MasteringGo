package main

import (
	"fmt"
	"time"
)

/*
We can use channels to synchronize execution across goroutines.
When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
*/
func worker(done chan bool) { // done channel is used to indicate if this function's work is done
	fmt.Print("Working...")
	time.Sleep(time.Second) // this is to say that some processing is happening
	fmt.Println("Done!")

	done <- true // send value to the channel to notify we're done
}

func main() {
	done := make(chan bool, 1) // make a buffered channel
	go worker(done)            // start the goroutine
	<-done                     // This is a blocking receive. This will block until we receive a notification from the worker on the channel
	// removing the <- done will exit the program before worker is even started
}

/*
Working...Done!
*/
