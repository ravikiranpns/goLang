package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
)

type wordsCount struct {
	word  string
	count int
}

func (p wordsCount) String() string {
	return fmt.Sprintf("%s %d", p.word, p.count)
}

func regexWordCount() {
	fmt.Println("regexWordCount() : start")
	fileName := "main.go"

	regEx := regexp.MustCompile("[a-zA-Z']+")
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	text := string(bs)
	matches := regEx.FindAllString(text, -1)

	words := make(map[string]int)

	for _, match := range matches {
		words[match]++
	}

	var wordCounts []wordsCount
	for k, v := range words {
		wordCounts = append(wordCounts, wordsCount{k, v})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count > wordCounts[j].count
	})

	for i := 0; i < 10; i++ {
		fmt.Println(wordCounts[i])
	}
}
