package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	arr := []rune(s)

	for i < j {
		left := unicode.ToLower(arr[i])
		right := unicode.ToLower(arr[j])

		if !isLetterOrDigit(left) {
			i++
			continue
		}

		if !isLetterOrDigit(right) {
			j--
			continue
		}

		if left != right {
			return false
		}

		i++
		j--
	}

	return true
}

func isLetterOrDigit(s rune) bool {
	return unicode.IsLetter(s)
}

func main() {
	str := ""
	fmt.Println(isPalindrome(str))
	str = "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(str))
	str = "race a car"
	fmt.Println(isPalindrome(str))
}
