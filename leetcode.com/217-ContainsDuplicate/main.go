package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	defer duration(track("main"))
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
	fmt.Println(containsDuplicateV1(nums))
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
