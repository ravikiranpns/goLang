/*
WaitGroups
https://gobyexample.com/waitgroups

- Wait for multiple goroutines to finish, use a wait group.
- worker function to run in every goroutine.
- Sleep to simulate an expensive task
	func worker(id int) {
		fmt.Printf("Worker %d starting\n", id)

		time.Sleep(time.Second)
		fmt.Printf("Worker %d done\n", id)
	}

- WaitGroup is used to wail for all the goroutines launced to finish
- NOTE : if a WaitGroup is explicityl passed into functions, it
should be done by pointer.
	var wg sync.WaitGroup

- Launch several goroutines and increment the WaitGroup counter
for each
	for i := 1; i <= 5; i++ {
		wg.Add(1)
	}

- Avoid re-use of the same i value in each goroutine closure.
	i := i

- Wrap the worker call in a closure that makes sure to tell the
WaitGroup that this worker is done. This way the worker itself does
not have to be aware of the concurrency primitives involved in
its execution.
	go func() {
		defer wg.Done()
		worker(i)
	}()

- Block until the WaitGroup counter goes back to 0; all the workers
notified they're done.
	wg.Wait()
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
