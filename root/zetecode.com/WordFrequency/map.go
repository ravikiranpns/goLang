package main

import (
	"fmt"
	"strings"
)

func mapWordCount() {
	//fmt.Println("Word Frequency")

	str := "jai shree ram load balancer distributes load across different servers ram shree ram jai shree ram"
	fmt.Println("str : ", str)

	fields := strings.FieldsFunc(str, func(r rune) bool {
		return !('a' <= r && r <= 'z' || r == '\'')
	})

	fmt.Println("fields : ", fields)

	wordsCount := make(map[string]int)
	fmt.Println("make map - wordCount : ", wordsCount)

	for _, field := range fields {
		wordsCount[field]++
	}
	fmt.Println("wordCount : ", wordsCount)

	keys := make([]string, 0, len(wordsCount))
	fmt.Println("make - keys : ", keys, len(keys), cap(keys))

	for key := range wordsCount {
		keys = append(keys, key)
	}
	fmt.Println("keys : ", keys, len(keys), cap(keys))

	for _, key := range keys {
		fmt.Printf("%s %d\n", key, wordsCount[key])
	}
}
