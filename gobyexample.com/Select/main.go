/*

Select -
- Lets wait on multiple channel operations.
- The combination of goroutines and channels with select
is a powerful feature
- Select accross two channels
	c1 := make(chan string)
	c2 := make(chan string)

- Each channel will receive a value after some amount of time,
to simulate e.g blocking RPC operations executing in concurrnet
goroutines

go func () {
	time.Sleep(1 * time.Second)
	c1 <- "ONE"
}()

go func () {
	time.Sleep(1 * time.Second)
	c2 <- "TWO"
}()

- Use select to await both of these values simultaneously,
printing each one as it arrives.
for i := 0; i < 2; i++ {
	select {
	case m1 := <- c1:
		fmt.Println("Recvd :", m1)
	case m2 := <- c2:
		fmt.Println("Recvd :", m2)
	}
}
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "ONE"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "TWO"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println("Recvd :", m1)
		case m2 := <-c2:
			fmt.Println("Recvd :", m2)
		}
	}
}
