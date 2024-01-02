// Take a string representation of an integer,
// such as "12345", and insert commas
// every 3 places, as in "12,345".

// int version

package main

import (
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
func main() {
	fmt.Println("in main func")

	str := comma("12345")
	fmt.Println(str)
}
