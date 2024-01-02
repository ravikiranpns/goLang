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

func findDuplicates(arr []int) []int {
	seen := map[int]bool{}
	var dups []int

	for i := 0; i < len(arr); i++ {
		if seen[arr[i]] {
			dups = append(dups, i)
		}
		seen[arr[i]] = true
	}
	return dups
}

func main() {

	arr := []int{1, 2, 3, 4, 1, 1}
	arr1 := []int{1, 2, 3, 3, 4, 5, 5, 6}
	fmt.Println(containsDup(arr))
	//fmt.Println(containsDup([]int{1}))
	fmt.Println(findDuplicates(arr))
	fmt.Println(findDuplicates(arr1))
}
