package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
The primary mechanism for managing state in Go is communication over channels.
We saw this for example with worker pools.
There are a few other options for managing state though.
Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
*/

func main() {
	var ops uint64 // We’ll use an unsigned integer to represent our (always-positive) counter.

	var wg sync.WaitGroup // The sync package provides a WaitGroup to help us wait for all our goroutines to finish.

	for i := 0; i < 50; i++ { // Here we’ll spin up 50 goroutines that each increment the counter 1000 times.
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// if non-atomic ops++ is used to increment the counter, we’d likely get a different number, changing between runs, because the goroutines would interfere with each other.
				//ops++
				atomic.AddUint64(&ops, 1) // To atomically increment the counter we use AddUint64, giving it the memory address of our ops counter with the & syntax.
				// fmt.Println(atomic.LoadUint64(&ops)) // Reading atomics safely while they are being updated is also possible, using functions like atomic.LoadUint64.
			}
			wg.Done()
		}()
	}

	wg.Wait() // wait until all goroutines are done

	// It’s safe to access ops now because we know no other goroutine is writing to it.

	fmt.Println("ops: ", ops)
}

/*
$ go run src/go-by-example/39atomic-counters.go
*/
