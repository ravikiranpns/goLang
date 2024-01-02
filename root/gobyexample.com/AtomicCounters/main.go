/*

Atomic Counters
https://gobyexample.com/atomic-counters

- The primary mechanism for managing state is communication over
channels.
- Worker pools is one way
- Few other options for managing state though.
- To use sync/atomic package for atomic counters accessed by
multiple goroutines.

- Use an unsigned integer to represent our (always-positive)
counter
	var ops uint64

- A WaitGroup will help wait for all goroutines to finish
their work.
	var wg sync.WaitGroup

- Start 50 goroutines that each increment the counter
exactly 1000 times.
- To atomically increment the counter we use AddUint64, giving
it the memory address of our ops counter with the & syntax.

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

- Wait until all the goroutines are done.
	wg.Wait()

- It's safe to access ops, because no ohter goroutine is writing
to it. Reading atomics safely while they are being updated is
also possible, using functions like atomic.LoadUint64.

	fmt.Println("ops: ", ops)

*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("ops:", ops)
}
