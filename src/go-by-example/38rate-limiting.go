package main

import (
	"fmt"
	"time"
)

/*
Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.
*/

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond) // This limiter channel will receive a value every 200 milliseconds.

	for req := range requests {
		<-limiter
		fmt.Println("request==", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3) // burstyLimiter channel will allow bursts of up to 3 events.

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(200 * time.Millisecond) { // Every 200 milliseconds we’ll try to add a new value to burstyLimiter, up to its limit of 3.
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5) // simulate 5 more requests
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request**", req, time.Now())
	}
}

/*
request== 1 2022-02-28 21:50:17.228985 +0900 JST m=+0.201285126
request== 2 2022-02-28 21:50:17.428788 +0900 JST m=+0.401086168
request== 3 2022-02-28 21:50:17.628859 +0900 JST m=+0.601154751
request== 4 2022-02-28 21:50:17.828855 +0900 JST m=+0.801147376
request== 5 2022-02-28 21:50:18.028931 +0900 JST m=+1.001220959


// For the second batch of requests we serve the first 3 immediately because of the burstable rate limiting,
// then serve the remaining 2 with ~200ms delays each.these 3 requests will be processed immediately as a brust
request** 1 2022-02-28 21:50:18.029321 +0900 JST m=+1.001611626
request** 2 2022-02-28 21:50:18.029334 +0900 JST m=+1.001623959
request** 3 2022-02-28 21:50:18.02934 +0900 JST m=+1.001630001

request** 4 2022-02-28 21:50:18.229454 +0900 JST m=+1.201741334
request** 5 2022-02-28 21:50:18.429485 +0900 JST m=+1.401769918
*/
