package main

import (
	"fmt"
	"strconv"
)

func atoi(s string) int {
	var output int

	sign := "positive"

	if string(s[0]) == "-" {
		sign = "negtive"
		s = s[1:]
	} else if string(s[0]) == "+" {
		s = s[1:]
	}

	strLen := len(s)

	for i := 0; i < strLen; i++ {
		tempNum, _ := strconv.Atoi(string(s[i]))
		output = output*10 + tempNum
	}

	if sign == "negtive" {
		output = output * -1
	}

	return output
}

func main() {
	output := atoi("121")
	fmt.Println(output)
}
