/*

https://gobyexample.com/non-blocking-channel-operations
Non-Blocking Channel Operations
-Basic sends and receives on channels are blocking.
-However, use select with a default clause to implement
non-blocking sends, receives and even non-blocking multi-way selects.

	msgs := make(chan string)
	signls := make(chan bool)

- Non-Blocking receive
- If a value is available on messages than select will take the <-msgs case with
that value.
- If not it will immediately take the default case.
	select {
	case msg := <-msgs:
		fmt.Println("Recvd msg", msg)
	default:
		fmt.Println("No msg recvd")
	}

- Non-Blocking Send
- msg cannot be sent to the msgs channel, because the channel has no buffer
& there is no receiver. Hence the default case is selected

	msg := "Hi"
	select {
	case msgs <- msg:
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("No message sent")
	}

- Non-Blocking multi-way selects.
- Multiple cases above the default clause. Attempt non-blocking receives on both
msgs and singnals.
	select {
	case msg := <-msgs:
		fmt.Println("Recvd Msg", msg)
	case sig := <-signls:
		fmt.Println("Recvd Signal", sig)
	default:
		fmt.Println("No activity")
	}

*/
package main

import "fmt"

func main() {
	msgs := make(chan string)
	signls := make(chan bool)

	select {
	case msg := <-msgs:
		fmt.Println("Recvd msg", msg)
	default:
		fmt.Println("No msg recvd")
	}

	msg := "Hi"
	select {
	case msgs <- msg:
		fmt.Println("Sent message", msg)
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-msgs:
		fmt.Println("Recvd Msg", msg)
	case sig := <-signls:
		fmt.Println("Recvd Signal", sig)
	default:
		fmt.Println("No activity")
	}
}
