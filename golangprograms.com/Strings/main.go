package main

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/text/transform"

	"golang.org/x/text/encoding/simplifiedchinese"
)

func main() {

	// A string ASCII and Unicode
	str := "a世界"

	fmt.Println("String: ", str)

	fmt.Println("Rune Code Points: ")
	for _, runeValue := range str {
		fmt.Printf("%c: %U\n", runeValue, runeValue)
	}

	fmt.Println("UTF-8 Bytes: ")
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c: %x\n", str[i], str[i])
	}

	for _, char := range str {
		fmt.Printf("U+%04X ", char)
	}

	byteLen := len(str)
	runeLen := utf8.RuneCountInString(str)
	fmt.Println(byteLen)
	fmt.Println(runeLen)

	runes := []rune(str)
	fmt.Println(runes)

	encoder := simplifiedchinese.GB18030.NewEncoder()
	encoded, _, _ := transform.String(encoder, str)
	fmt.Println(encoded)
}
