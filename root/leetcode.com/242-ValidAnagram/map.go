package main

//https://golangbyexample.com/check-two-strings-anagram-go/

func isAnagramMap(s string, t string) bool {
	lenS := len(s)
	lenT := len(t)

	if lenS != lenT {
		return false
	}

	m := make(map[string]int)

	for i := 0; i < lenS; i++ {
		m[string(s[i])]++
	}

	for i := 0; i < lenT; i++ {
		m[string(t[i])]--
	}

	for i := 0; i < lenS; i++ {
		if m[string(s[i])] != 0 {
			return false
		}
	}
	return true
}
