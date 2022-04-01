// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Hello, 世界")
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	go PrintTick(ctx)

	time.Sleep(5 * time.Second)
	cancel()

	wg.Wait()
}

func PrintTick(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done: ", ctx.Err())
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}
