/*
https://gobyexample.com/tickers
Tickers

- To do something repeatedly at regular intervals.
- Example of a ticker, ticks periodically until stoped it.
- Tickers use a similar mechanism to timers: a channel that
is sent values.
- The select builtin on the channel to await the values as they
arrive every 500ms.

	ticker := time.NewTicker(500 * time.Millicsecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

- Tickers can be stopped like timers.
- Once a ticker is stopped it won't receive any more values
on its channel.
- Stop after 1600ms
		time.Sleep(1600 * time.Millicsecond)
		ticker.Stop()
		done <- true
		fmt.Println("Ticker stopped")
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
