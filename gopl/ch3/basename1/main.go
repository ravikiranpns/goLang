package main

import "fmt"

//basename removes directory components and a .suffix.
// eg a > a, a.go > a, a/b/c.go > c, a/b.c.go > b.c
func basename(s string) string {
	// Discard last	'/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("abc.go"))

}
