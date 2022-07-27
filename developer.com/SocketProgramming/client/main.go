package main

import (
	"fmt"
	"net"
)

/*
Intro to Socket Programming in Go

https://www.developer.com/languages/intro-socket-programming-go/

*/

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)

	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("Hello Server! Greetings."))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
	}
	buffer := make([]byte, 1024)

	mLen, err := conn.Read(buffer)

	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	fmt.Println("Received: ", string(buffer[:mLen]))

	defer conn.Close()
}
