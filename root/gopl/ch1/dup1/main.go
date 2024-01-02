package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("main() : start dup1")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	fmt.Println("main() : end dup1")
}
