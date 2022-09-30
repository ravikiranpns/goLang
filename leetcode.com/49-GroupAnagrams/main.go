package main

import "fmt"

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagramsBf(strs))
	fmt.Println(groupAnagramsEff(strs))
}
