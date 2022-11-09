/*
https://gobyexample.com/closing-channels

Closing Channels -
- On closing channels - no more values will be sent on it.
- Useful to communicate completion to the channels receivers
- Jobs Channel
	- to communicate work to be done from the main() goroutine to a
	worker goroutine
	- when no more jobs for the worker, close the jobs channel

	jobs := make(chan int, 5)
	defer close(jobs)

- Worker goroutine
	- it repeateadly receives from jobs with j, more := <-jobs.
	- the special 2-value form of receive, the more value will be
	false if jobs has been closed and all values in the channel
	have already been received.
	- Use this to notify on done when worked all jobs.

	done := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Recvd job", j)
			} else {
				fmt.Println("Recvd all jobs")
				done <- true
				return
			}
		}
	}()

- To send 3 jobs to the worker over the jobs channel, then close it
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent job", j)
	}
	close(jobs)
	fmt.Println("Sent all jobs")

- await the worker using the synchronization approach
	<-done

*/

package main

import "fmt"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("Recvd job", j)
			} else {
				fmt.Println("Recvd all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent job", j)
	}
	close(jobs)
	fmt.Println("Sent all jobs")

	<-done
}
