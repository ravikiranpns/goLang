package main

import "fmt"

func groupAnagram(strs []string) [][]string {
	am := make(map[[26]int][]string)

	for _, str := range strs {
		var freq [26]int

		for _, c := range str {
			freq[c-'a']++
		}
		am[freq] = append(am[freq], str)
	}
	result := make([][]string, len(am))

	i := 0

	for _, v := range am {
		result[i] = v
		i++
	}
	return result
}

func main() {

	strs := []string{"ra", "ar", "ka", "ak"}
	fmt.Println(groupAnagram(strs))
}
