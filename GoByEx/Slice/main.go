package main

import "fmt"

/*
1686400.22.06.22
https://golangbyexample.com/slice-in-golang/

Internal Representation
- Pointer to the underlying array
- Current length of the Underlying array
- Total Capacity which is the maximum capacity to which the underlying array can expand.

type SliceHeader strcut {
	Pointer uintptr
	Len 	int
	Cap		int
}

Creating
- Using the []<type>{}format
	s := []int --- empty slice of 0 len and 0 Cap
	s := []int{1, 2} --- len 2, cap 2
	len() function - for length of the slice
	cap() func -	for capacity of the slice
- Creating from another slice or array
	[n]sample[start:end]
- Using make
	func make([]{type}, length, capacity int) []{type}
- Using new

*/

func main() {
	sample := []int{}

	fmt.Println(len(sample))
	fmt.Println(cap(sample))
	fmt.Println(sample)

	letters := []string{"a", "b", "c"}
	fmt.Println(len(letters))
	fmt.Println(cap(letters))
	fmt.Println(letters)

	numbers := [5]int{1, 2, 3, 4, 5}
	num1 := numbers[2:4]
	fmt.Printf("num1= %v\n", num1)
	fmt.Printf("len = %d\n", len(num1))
	fmt.Printf("cap = %d\n", cap(num1))
	num1 = numbers[2:]
	fmt.Printf("num1= %v\n", num1)
	fmt.Printf("len = %d\n", len(num1))
	fmt.Printf("cap = %d\n", cap(num1))

	num1 = numbers[:3]
	fmt.Printf("num1= %v\n", num1)
	fmt.Printf("len = %d\n", len(num1))
	fmt.Printf("cap = %d\n", cap(num1))

	num1 = numbers[:]
	fmt.Printf("num1= %v\n", num1)
	fmt.Printf("len = %d\n", len(num1))
	fmt.Printf("cap = %d\n", cap(num1))

	num := make([]int, 3, 5)
	fmt.Printf("make num = %v\n", num)
	fmt.Printf("length=%d\n", len(num))
	fmt.Printf("capacity=%d\n", cap(num))

	num = make([]int, 3)
	fmt.Printf("make num = %v\n", num)
	fmt.Printf("length=%d\n", len(num))
	fmt.Printf("capacity=%d\n", cap(num))

	numbNew := new([]int)
	fmt.Printf("numbNew=%v\n", *numbNew)
	fmt.Printf("Len = %d\n", len(*numbNew))
	fmt.Printf("Capacity = %d\n", cap(*numbNew))
}
