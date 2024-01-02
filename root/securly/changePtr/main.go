// https://www.includehelp.com/golang/pointer-programs.aspx
package main

import "fmt"

func main() {
	v := 5
	p := &v

	var num int = 10
	var ptr *int = &num
	fmt.Printf("num: %d\n", num)
	fmt.Printf("*ptr: %d\n", *ptr)
	changePtr(ptr)
	fmt.Printf("After changePtr: %d\n", *ptr)

	fmt.Println(*p)
	changePtr(p)
	fmt.Println(*p)
}

func changePtr(p *int) {
	v := 3
	p = &v
	fmt.Println(*p)
}
