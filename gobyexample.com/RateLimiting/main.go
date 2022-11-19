/*
https://gobyexample.com/rate-limiting

Rate Limiting
-	Mechanism for controlling resource utilization and maintaining
quality of service.
- Elegantly support rate limiting with goroutines, channels,
and tickers.

- Basic Rate Limiting
	- To limit handling of incomming requests.
	- Serve the requests off a channel of the same name.
		requests := make(chan int, 5)
		for i := 1; i <= 5; i++ {
			requests <- i
		}
		close(requests)

- The limiter channel will receive a value every 200 milliseconds.
- The regulator in rate limiting scheme.
	limiter := time.Tick(200 * time.Millisecond)

- By blocking on a receive from the limiter channel before serving
each request, limit to 1 request every 200 milliseconds.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

- To allow short bursts of requests in rate limiting scheme
while preserving the overall rate limit.
- Accomplish this by buffering limiter channel. The burstyLimiter
channel will allow burst of up to 3 events.
	burstyLimiter := make(chan time.Time, 3)

- Fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

- Every 200 milliseconds try to add a new value to burstyLimiter,
up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

- Simulate 5 more incomming requests. The first 3 of these will
benefit from the burst capability of burstyLimiter.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

*/

package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
