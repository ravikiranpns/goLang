package main

import (
	"fmt"
	"time"
)

func main() {

	//ch := make(chan<- struct{}{})
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d ", i)
			time.Sleep(1 * time.Second)

		}
	}()

	go func() {
		for i := 'a'; i <= 'j'; i++ {
			fmt.Printf("%c ", i)
			time.Sleep(1 * time.Second)
		}
	}()

	//<-chan struct{}
}
