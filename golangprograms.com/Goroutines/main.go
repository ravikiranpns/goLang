package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func responseSize(url string, size chan int) {

	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	size <- len(body)
}

func main() {
	size := make(chan int)

	wg.Add(1)
	go responseSize("https://www.golangprograms.com", size)
	fmt.Println(<-size)
	wg.Wait()
	close(size)
}
