package main

import (
	"fmt"
	"sync"
	"time"
)

/*
To wait for multiple goroutines to finish, we can use a wait group.
*/

func workerFunction(id int) { // This is the function we’ll run in every goroutine.
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // indicates that this is an expensive task
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// This WaitGroup is used to wait for all the goroutines launched here to finish.
	// Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // increment the WaitGroup counter for each i

		i := i // Avoid re-use of the same i value in each goroutine closure. See https://go.dev/doc/faq#closures_and_goroutines for more details.

		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
		// This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
		go func() { // Launch several goroutines. In this case it's 5 goroutines
			defer wg.Done() // Done decrements the WaitGroup counter by one.
			workerFunction(i)
		}()
	}

	wg.Wait() // Block until the WaitGroup counter goes back to 0; all the workers notified they’re done.
}

// Despite having 5 worker which takes 1 second each, the whole thing completes in 1.541 seconds.
/*
Worker 5 starting
Worker 2 starting
Worker 3 starting
Worker 1 starting
Worker 4 starting
Worker 4 done
Worker 1 done
Worker 5 done
Worker 2 done
Worker 3 done
go run 37waitgroups.go  0.13s user 0.16s system 18% cpu 1.541 total
*/
