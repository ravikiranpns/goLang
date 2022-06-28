package main

import "fmt"

func main() {
	arr := []int{90, 70, 30, 10, 80, 77, 70, 77, 77}

	freq := make(map[int]int)
	for _, num := range arr {
		freq[num] = freq[num] + 1
	}

	freq_map := freq_map(arr)
	fmt.Println("Frequency of the Array using func : ", freq_map)
	fmt.Println("Frequency of the Array : ", freq)
}

func freq_map(arr []int) map[int]int {
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num] = freq[num] + 1
	}
	return freq
}
