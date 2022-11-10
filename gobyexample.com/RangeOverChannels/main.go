/*

https://gobyexample.com/range-over-channels

Range Over Channels
- The for and range provide iteration over basic data structures.
- The syntax can be used to iterate over values received from a channel
- Iterate over 2 values in the queue channel
	queue := make(chan string, 2)
	queue <- "One"
	queue <- "two"
	close(queue)
- The range iterates over each element as it's received from queue.
- Because channel is closed above,the iteration terminates after receiving
the 2 elments.
	for elem := range queue {
		fmt.Println(elem)
	}

*/

package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "One"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
