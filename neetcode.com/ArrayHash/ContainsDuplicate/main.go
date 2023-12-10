package main

import "fmt"

func containsDup(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	xm := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := xm[v]; ok {
			return true
		}
		xm[v] = struct{}{}
	}
	return false
}

func main() {

	arr := []int{1, 2, 3, 4, 1}
	fmt.Println(containsDup(arr))
}
