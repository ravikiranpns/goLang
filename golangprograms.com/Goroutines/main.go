package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
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
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://coderwall.com")
	time.Sleep(2 * time.Second)
}
