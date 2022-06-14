package main

import "fmt"

func main() {
	//var country map[int]string
	var country = make(map[int]string)

	country[1] = "Bharath"
	country[2] = "Nepal"
	country[3] = "Japan"
	country[4] = "Srilanka"
	country[5] = "Russia"
	country[6] = "Israil"
	country[7] = "USA"

	for i, j := range country {
		fmt.Printf("Key: %d, Value : %s\n", i, j)
	}
}
