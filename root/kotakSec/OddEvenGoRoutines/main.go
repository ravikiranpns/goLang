package main

import (
	"fmt"
	"time"
)

func evenNumbers(ch chan<- int) {
	for i := 2; i <= 100; i += 2 {
		ch <- i
		time.Sleep(2 * time.Millisecond)
	}
}

func oddNumbers(ch chan<- int) {
	for i := 1; i <= 100; i += 2 {
		ch <- i
		time.Sleep(2 * time.Millisecond)
	}
}

func main() {
	evenChan := make(chan int)
	oddChan := make(chan int)

	go evenNumbers(evenChan)
	go oddNumbers(oddChan)

	for i := 0; i < 10; i++ {
		select {
		case num := <-evenChan:
			fmt.Printf("Even: %d\n", num)
		case num := <-oddChan:
			fmt.Printf("Odd: %d\n", num)
		}
	}
}
