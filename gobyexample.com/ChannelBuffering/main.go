/*
Channel Buffering
- By default channels are unbuffered, that they will only accept sends
(chan <-) if there is a corresponding receive (<-chan) ready to receive
the sent value.
	sends (chan <-)
	receive (<-chan)

- Buffered channels accept a limited number of values without a correspo
-nding receiver for those values.

- Make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

- The channel of strings buffered, send these values into the channel
without a corresponding concurrent receive.
	messages <- "buffered"
	messages <- "channel"

- Later receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
*/

package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Printf("%T\n", messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
