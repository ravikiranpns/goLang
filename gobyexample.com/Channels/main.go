/*
https://gobyexample.com/channels

Channels
- The Pipes that connect concurrent goroutines.
- Send values into channels from one goroutine and
receive those values into another gorotine.
- Create a new channel - make(chan val-type). Channels
are typed by the values they convey.
	messages := make(chan string)

- Send a value into a channel using the channel <- syntax.
- Send "ping" to the message channel from a new goroutine
	go func() { messages <- "ping" } ()

- The <-channel syntax receives a value from the channel.
- Receive the "ping" message
	msg := <-messages

- While running the "ping" message is successfully passed from one
goroutine to another via channel.

- By default sends and receives block until both the sender and
receiver are ready. This propery allowed to wait at the end of
program for the "ping" message without having to use any other
synchronization.

*/

package main

import "fmt"

func main() {
	messages := make(chan string)
	count := make(chan int)
	fmt.Printf("%v, %T\n", messages, messages)

	go func() { messages <- "ping" }()

	go func() { count <- 1 }()

	msg := <-messages

	fmt.Println(<-count)
	fmt.Println(messages)
	fmt.Println(msg)
}
