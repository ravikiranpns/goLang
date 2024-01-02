/*

Channel Synchronization
- Channels used to synchronize execution across goroutines.
- Using Blocking receive to wait for a goroutine to finish.
- When waiting for multiple goroutines to finish, WaitGroup is used.

- The worker func will run in a goroutine. The done channel will be used
to notify another goroutine that this funtion's work is done.
	func worker(done chan bool) {
		fmt.Print("Working...")
		time.Sleep(time.Second)
		fmt.Println("done")

		done <- true
	}

- Start a worker goroutine, giving it the channel to notify on.
	func main() {
		done := make(chan bool, 1)
		go worker(done)

		<- done
	}
*/

package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
