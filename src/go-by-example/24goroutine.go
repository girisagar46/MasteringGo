package main

import (
	"fmt"
	"time"
)

/*
A goroutine is a lightweight thread of execution.
*/

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct call") // this is a synchronous function call

	go f("goroutine call") // this function will execute concurrently

	go func(msg string) { // goroutine is also possible in anonymous function
		fmt.Println(msg)
	}("goroutine anon func call")

	time.Sleep(time.Second)

	fmt.Println("DONE!")
}

/*
we see the output of the blocking call first, then the output of the two goroutines.
The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.
*/

/*
direct call : 0
direct call : 1
direct call : 2
goroutine call : 0
goroutine call : 1
goroutine anon func call
goroutine call : 2
DONE!
*/
