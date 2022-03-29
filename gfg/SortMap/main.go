package main

import (
	"fmt"
	"sort"
)

func main() {
	basket := map[string]int{"orange": 5, "apple": 7, "mango": 3}

	keys := make([]string, 0, len(basket))

	for k := range basket {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, basket[k])
	}
}
