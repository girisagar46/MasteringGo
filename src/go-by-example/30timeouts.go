package main

import (
	"fmt"
	"time"
)

/*
Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
*/

func main() {
	// the channel c1 is buffered, so the send in the goroutine is nonblocking.
	// This is a common pattern to prevent goroutine leaks in case the channel is never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second) // say we're requesting a URL that takes 2 seconds
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // `<-time.After` awaits a value to be sent after the timeout of 1s.
		// But the result arrives in 2 sec (see line 13). Hence, timeout happens and this case is matched.
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second) // get result in 2 seconds
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second): // but timeout is of 3 seconds. Hence, timeout does not happen in this case.
		fmt.Println("timeout 2")
	}
}

/*
timeout 1
result 2
*/
