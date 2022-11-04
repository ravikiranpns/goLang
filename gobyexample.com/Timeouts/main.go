/*
Timeouts
- Important for programs that connect to external resources or that
otherwise need to bound execution time.
- Implementing timeouts in Go is easy and elegant(using channels & select)
- Suppose executing an external call that returns its result on a channel c1
after 2s
- The channel is buffered, so the send in the goroutines in non blocking.
- This is a common pattern to prevent goroutines leaks in case the channel is
never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Result 1"
	}()

- The Select implementing a timeout. res := <- c1 awaits the result and
	<-time.After awaits a value to be sent after the timeout of 1s

- Since select proceeds with the first receive that's ready to take the
timeout case if the operation takes more than the allowed 1s.

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}
- If we allow a longer time out of 3s, then the receive from c2 will succeed &
print the result
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout 2")
	}
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout 2")
	}
}
