/*
https://gobyexample.com/timers
Timers
- To execute the code at some point in the future, or repeateadly at some interval.
- Built-in timer and ticker features make both of these tasks easy.
- Timers represent a single event in the future.
- Tell timer how long to wait and it provides a channel that will be notified
at that time
	timer1 := time.NewTimer(2 * time.Second)

- The <-timer1.C blocks on the timer's channel C until it sends a value. Indicating that
the timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

- To wait, time.Sleep. One reason a timer may be useful is to cancel the timer before
it fires.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<- timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)


*/

package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
