package main

func isAnagramCount(s string, t string) bool {

	sCount := countAlphabet(s)
	tCount := countAlphabet(t)

	for i := 0; i < 26; i++ {
		if sCount[i] != tCount[i] {
			return false
		}
	}
	return true
}

func countAlphabet(s string) []int {
	count := make([]int, 26)
	for alphabet := range s {
		count[s[alphabet]-'a']++
	}

	return count
}
