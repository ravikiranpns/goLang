package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var freq [26]int

	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for i := 0; i < len(s); i++ {
		if freq[i] != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("ravi", "kiran"))
	fmt.Println(isAnagram("ravi", "rvai"))
	fmt.Println(isAnagram("abc", "cca"))

}
