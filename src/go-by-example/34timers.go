package main

import (
	"fmt"
	"time"
)

/*
Go’s built-in timer and ticker features make following tasks easy:
1. execute Go code at some point in the future,
2. execute Go code repeatedly at some interval
*/
func main() {
	/*
		Timers represent a single event in the future.
		You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.
	*/
	fmt.Println(time.Now(), "program start")
	timer1 := time.NewTimer(2 * time.Second) // this timer waits for 2 seconds

	<-timer1.C // this blocks on the timer’s channel C until it sends a value indicating that the timer fired.
	// we could also use time.Sleep if we just wanted to wait
	fmt.Println(time.Now(), "Timer1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println(time.Now(), "Timer2 fired")
	}()

	stop2 := timer2.Stop() // we can cancel the timer before it fires.
	if stop2 {
		fmt.Println(time.Now(), "Timer2 stopped.")
	}

	time.Sleep(2 * time.Second) // the wait is going to happen for 2 seconds, but we already Stop the timer.
}

// The first timer will fire ~2s after we start the program, but the second should be stopped before it has a chance to fire
/*
2022-02-28 09:16:00.840348 +0900 JST m=+0.000126251 program start
2022-02-28 09:16:02.841705 +0900 JST m=+2.001486709 Timer1 fired
2022-02-28 09:16:02.841853 +0900 JST m=+2.001634959 Timer2 stopped.
*/
