package main

import "fmt"

func topKFreqElements(nums []int, k int) (res []int) {
	cm := map[int]int{}
	cs := make([][]int, len(nums)+1)

	// Add key - value(num - count) in the map
	for _, num := range nums {
		if count, ok := cm[num]; ok {
			cm[num] = count + 1
		} else {
			cm[num] = 1
		}
	}

	// append slice
	for num, count := range cm {
		cs[count] = append(cs[count], num)
	}

	for i := len(cs) - 1; i > 0; i-- {
		res = append(res, cs[i]...)
		if len(res) == k {
			return
		}
	}
	return
}

func main() {
	nums := []int{1, 1, 2, 3, 4, 3, 4, 3, 3, 2, 2}
	fmt.Println(topKFreqElements(nums, 2))
}
