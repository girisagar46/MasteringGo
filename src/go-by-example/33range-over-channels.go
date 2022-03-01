package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	close(queue) // even after closing the channel, we can receive the value from channel

	for elem := range queue { // iterate over values received from the channel queue
		// Because we closed the channel above, the iteration terminates after receiving the 2 elements.
		fmt.Println(elem)
	}
}
