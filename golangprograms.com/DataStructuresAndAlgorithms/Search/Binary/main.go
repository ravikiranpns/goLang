package main

import "fmt"

// binarySearch func return true if key to be searched is present else returns false.
func binarySearch(key int, slc []int) bool {
	low := 0
	high := len(slc) - 1

	for low <= high {
		mid := (low + high) / 2

		if slc[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low == len(slc) || slc[low] != key {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 3, 5, 7, 8, 9, 10}
	fmt.Println(binarySearch(7, items))
}
