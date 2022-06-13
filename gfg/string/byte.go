// Go program to illustrate how to
// access the bytes of the string
package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a string
	str := "Welcome to GeeksforGeeks"

	// Accessing the bytes of the given string
	for c := 0; c < len(str); c++ {

		fmt.Printf("\nCharacter = %c Bytes = %v", c, str)
	}
}
