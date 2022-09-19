package main

import (
	"log"
	"time"
)

func containsDuplicate(nums []int) bool {

	/*
			for i := 0; i < len(nums); i++ {
			for j := 1; j < len(nums); j++ {
				if nums[i] == nums[j] {
					return true
				}
			}
		}
	*/
	start := time.Now()
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	elaplsed := time.Since(start)
	log.Printf("Brute Fore took %s", elaplsed)
	return false
}
