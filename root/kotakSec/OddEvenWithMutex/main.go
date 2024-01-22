package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex // mutex for mutual exclusion
var even bool     // flag to check
func evenNumbers(ch chan<- int) {
	for i := 2; i <= 10; i += 2 {
		mu.Lock()
		if even {
			fmt.Printf("Even: %d\n", i)
		} else {
			fmt.Printf("Odd: %d\n", i)
		}
		even = !even
		mu.Unlock()
		time.Sleep(1 * time.Millisecond)

	}
}

func oddNumbers(ch chan<- int) {
	for i := 1; i <= 10; i += 2 {
		mu.Lock()
		if !even {
			fmt.Printf("Even: %d\n", i)
		} else {
			fmt.Printf("Odd: %d\n", i)
		}

		even = !even
		mu.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	evenChan := make(chan int)
	oddChan := make(chan int)

	go evenNumbers(evenChan)
	go oddNumbers(oddChan)

	for i := 0; i < 100; i++ {
		select {
		case num := <-evenChan:
			fmt.Printf("Even: %d\n", num)
		case num := <-oddChan:
			fmt.Printf("Odd: %d\n", num)
		}
	}
}
