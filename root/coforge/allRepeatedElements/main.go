//Problem 1: Write a function which will accept the given array as argument and will return the repeating elements from the given array.
// Given array: []int{1, 2, 3, 4, 3, 2, 4, 6, 2, 9, 7, 4}
//Expected Output: [3 2 4]
// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 3, 2, 4, 6, 2, 9, 7, 4}
	fmt.Println("Elements are:", a)
	fmt.Println("Duplicate Elements are:", getDuplicateElements(a))
}

func getDuplicateElements(a []int) []int {
	// Your logic start here
	m := make(map[int]struct{})
	var res []int = []int{}
	//i := 0
	for _, v := range a {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
		m[v] = struct{}{}
	}
	return res
}
