package main

import "fmt"

func longestConsecutive(nums []int) int {
	set := map[int]bool{}

	for _, num := range nums {
		set[num] = true
	}

	res := 0

	for _, num := range nums {
		if set[num-1] {
			continue
		}

		sequence := 1
		temp := num + 1

		for set[temp] {
			sequence++
			temp++
		}

		if sequence > res {
			res = sequence
		}
	}

	return res
}

func main() {
	fmt.Println(longestConsecutive([]int{}))
	fmt.Println(longestConsecutive([]int{100}))

	fmt.Println(longestConsecutive([]int{5, 1, 4, 100}))
	fmt.Println(longestConsecutive([]int{5, 1, 4, 3, 100, 99, 98}))

	var arr [5]float32
	var sliceInt [5]int
	var r rune

	fmt.Println(len(arr), cap(arr), arr)
	fmt.Println(len(sliceInt), cap(sliceInt), sliceInt)
	fmt.Println(r)
	//len(r), cap(r),
}
