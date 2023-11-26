package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func responseSize(url string) {

	// Schedule the call to WaitGroup's Done to tell goroutine is completed
	defer wg.Done()

	fmt.Println("S1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("S2: ", url)
	defer response.Body.Close()

	fmt.Println("S3: ", url)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("S4: ", len(body))
}

func main() {

	// Add a count, one for each goroutine.

	wg.Add(2)
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	time.Sleep(2 * time.Second)

	wg.Wait()

	fmt.Println("Terminating main")
}
