package main

import (
	"fmt"
	"time"
)

/*
https://gobyexample.com/goroutines

1826100.10.08.22

Goroutines
- A lightweight thread of execution.

	func f(from string) {
		for i := 0; i < 3; i++ {
			fmt.Println(from, ":", i)
		}
	}

- function call f(s). call that in the usual way, running it synchronously.

	f("direct-synchronously")

- Invoke the function in a goroutine, use go f(s). This new goroutine will
execute concurrently with the calling one.

	go f("goroutine-concurrently-asynchronously")

  - Start a goroutine for an anonymous function call.
    go func(msg string) {
    fmt.Println(msg)
    }("anonymous goroutine-going")

- Two function calls are running asynchronously in separate goroutines now.
Wait for them to finish(for a more robust approch, Use WaitGroup).

	time.Sleep(time.Second)
	fmt.Println("done")
*/
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct-synchronously-Blocking Call")
	go f("goroutine-concurrently-asynchronously-nonBlocking")

	go func(msg string) {
		fmt.Println(msg)
	}("anonymous func goroutine-going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
