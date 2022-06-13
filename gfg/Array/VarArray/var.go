package main

import "fmt"

func main() {
	var arr [3]string

	arr[0] = "Om"
	arr[1] = "Shree"
	arr[2] = "Ram"

	fmt.Println("Elements of Arr")
	fmt.Println("1 : ", arr[0])
	fmt.Println("2 : ", arr[1])
	fmt.Println("3 : ", arr[2])
	fmt.Println("Length of arr : ", len(arr))
}
