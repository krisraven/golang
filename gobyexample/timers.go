// sometimes go code needs to be executed at some point, or at repeated intervals. Go's built-in timer and ticker features make both of these tasks easy. We'll first look at timers then at tickers

package main

import (
	"fmt"
	"time"
)

// timers represent a single event in the future. The timer is told how long to wait, and it provides a channel that will be notified at that time. Eg. This timer is 2 seconds
func main() {
	timer1 := time.NewTimer(2 * time.Second)
	// the timer1.C blocks on the timers channel C until it sends a value indicating that the timer fired
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// time.Sleep could have been used instead
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// timers can be stopped at some point before it fires. This is an exmaple of that
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	// the first timer will fire ~2secs after the program is started, but the second should not be stopped before it has a chance to fire
	time.Sleep(2 * time.Second)
}
