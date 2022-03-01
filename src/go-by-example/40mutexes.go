package main

import (
	"fmt"
	"sync"
)

/*
We use atomic operations to manage simple counter state.
We use mutext to safely access data across multiple goroutines
*/

//Container holds a map of counters; since we want to update it concurrently from multiple goroutines, we add a Mutex to synchronize access.
//Note that mutexes must not be copied, so if this struct is passed around, it should be done by pointer.
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) Inc(name string) {
	c.mu.Lock()         // Lock the mutex before accessing counters; this ensures that the map is not being updated while we are reading it.
	defer c.mu.Unlock() // Unlock the mutex when we are done reading from counters.

	c.counters[name]++
}

func main() {
	c := Container{ // Note that the zero value of a mutex is usable as-is, so no initialization is required here.
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) { // This function increments a named counter in a loop.
		for i := 0; i < n; i++ {
			c.Inc(name)
		}
		wg.Done()
	}

	// Run several goroutines concurrently; note that they all access the same Container, and two of them access the same counter.
	wg.Add(3)
	go doIncrement("a", 100000)
	go doIncrement("a", 100000)
	go doIncrement("b", 100000)

	wg.Wait()

	fmt.Println(c.counters)
}

// The output is correct as expected because the mutex is locked during the reading of the map, and unlocked when we are done.
/*
map[a:200000 b:100000]
*/
