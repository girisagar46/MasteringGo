package main

import (
	"fmt"
	"time"
)

/*
tickers are for when you want to do something repeatedly at regular intervals.
*/
func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	/*
		Tickers use a similar mechanism to timers: a channel that is sent values.
		Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.
	*/
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C: // get values from ticker channel in regular interval. The value will be empty once the ticker is stopped
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop() // Once a ticker is stopped it won’t receive any more values on its channel.
	done <- true  // this is to break the select
	fmt.Println("Ticker stopped!")
}

// Every tick happened exactly after 500ms
/*
Tick at 2022-02-28 09:18:33.891152 +0900 JST m=+0.501102376
Tick at 2022-02-28 09:18:34.391184 +0900 JST m=+1.001134793
Tick at 2022-02-28 09:18:34.891042 +0900 JST m=+1.500993668
Ticker stopped!

*/
