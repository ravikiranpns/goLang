package main

// https://leetcode.com/problems/group-anagrams/

import (
	"strconv"
)

func groupAnagramsBf(strs []string) [][]string {
	var ans [][]string
	dict := make(map[string][]string)

	for _, str := range strs {
		tmpArr := [26]int{}

		for _, char := range str {
			tmpArr[char-'a']++
		}
		tmpStr := ""
		for _, val := range tmpArr {
			tmpStr += strconv.Itoa(val)
		}
		dict[tmpStr] = append(dict[tmpStr], str)
	}
	for _, val := range dict {
		ans = append(ans, val)
	}

	return ans
}
