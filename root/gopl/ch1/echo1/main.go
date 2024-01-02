// Echo1 prints command-line args.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

/*
index i = -1
for i := -1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

Apples-MacBook-Pro:echo1 apple$ go run main.go OM OM
panic: runtime error: index out of range [-1]

goroutine 1 [running]:
main.main()
        /Users/ravikiran/go/src/github.com/goLang/gopl/ch1/echo1/main.go:14 +0xd8
exit status 2

*/
