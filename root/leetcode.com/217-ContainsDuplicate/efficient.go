package main

import (
	"log"
	"time"
)

func containsDuplicateV1(nums []int) bool {
	start := time.Now()

	m := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}

	elapsed := time.Since(start)
	log.Printf("efficient took %s", elapsed)
	return false
}
