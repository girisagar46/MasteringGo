package main

import (
	"fmt"
	"time"
)

/*
worker pool will wait for multiple goroutines
An alternative way to wait for multiple goroutines is to use a WaitGroup.
*/

// These workerFunc will receive work on the jobs channel and send the corresponding results on results.
func workerFunc(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second) // We’ll sleep a second per job to simulate an expensive task.s
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)    // the jobs channel
	results := make(chan int, numJobs) // the results channel

	for w := 1; w <= 3; w++ { // create 3 instances of the workerFunc
		go workerFunc(w, jobs, results) // run 3 instances of the workerFunc in goroutine
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j // push 5 jobs into jobs channel
	}
	close(jobs) // close the jobs channel to indicate that’s all the work we have.

	for a := 1; a <= numJobs; a++ { // Finally, we collect all the results of the work. This also ensures that the worker goroutines have finished.
		<-results
	}
}

/*
Our running program shows the 5 jobs being executed by various workers.
The program only takes about 2 seconds despite doing about 5 seconds of total work because there are 3 workers operating concurrently.
*/

/*
$ time go run 36worker-pools.go
worker 1 started job 2
worker 3 started job 1
worker 2 started job 3
worker 2 finished job 3
worker 2 started job 4
worker 1 finished job 2
worker 1 started job 5
worker 3 finished job 1
worker 1 finished job 5
worker 2 finished job 4
go run 36worker-pools.go  0.13s user 0.16s system 11% cpu 2.523 total
*/
