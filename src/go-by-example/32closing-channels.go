package main

import "fmt"

/*
Closing a channel indicates that no more values will be sent on it.
This can be useful to communicate completion to the channel’s receivers.
*/
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// this is a special 2-value form of receive, the `more` value will be false if jobs has been closed and all values in the channel have already been received.
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else { // when jobs channel is closed
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ { // sends 3 jobs to the worker over the jobs channel, then closes it. (closed at line 32)
		fmt.Println("sending job...", j)
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}

/*
sending job... 1
sent job 1
sending job... 2
sent job 2
sending job... 3
sent job 3
sent all jobs
received job 1
received job 2
received job 3
received all jobs
*/
