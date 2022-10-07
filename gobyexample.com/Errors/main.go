package main

import (
	"errors"
	"fmt"
)

/*
https://gobyexample.com/errors
1824500.07.10.22

Errors
- It's idiomatic to communicate errors via an explicit, seprate return value.
- The approach makes it easy to see which functions return errors and to
handle them using the same language constructs employed for any other,
non-error tasks.

- Pkg - "errors"\

- By convention, errors are the last return value and have type error,
a built-in interface

- errors.New constructs a basic error value with the given error message.
- A nil value in the error position indicates that there was no error.

	func f1(arg int) (int, error) {
		if arg == 42 {
			return -1, errors.New("can't work with 42")
		}
		return arg + 3, nil
	}

- Possible to use custom types as errors by implementing the Error() method
on them. Variant on the example above that uses a custom type to explicitly
represent an argument error.

	type argError struct {
		arg int
		prob string
	}

	func (e *argError) Error() string {
		return fmt.Sprintf("%d - %s", e.arg, e.prob)
	}

- Use &argError syntax to build a new struct, supplying values for the
two fields arg and prog.

	func f2(arg int) (int, error) {
		if arg = 42 {
			return -1, &argError{arg, "can't work with it"}
		}
		return arg + 3, nil
	}

- The 2 loops to test out each of error-returning functions.
Note - the use of an inline error check on the if line is a common
idiom

- To Programmatically use the data in a custom error, need to get the error
as an instance of the custom error type via type assertion.

	func main() {
		for _, i := range []int{7, 42} {
			if r, e := f1(i); e != nil {
				fmt.Println("f1 failed:", e)
			} else {
				fmt.Println("f1 worked:", r)
			}
		}
		for _, i := range []int{7, 42} {
			if r, e := f2(i); e != nil {
				fmt.Println("f2 failed:", e)
			} else {
				fmt.Println("f2 worked:", r)
			}
		}

		_, e := f2(42)
		if ae, ok := e.(*argError); ok {
			fmt.Println(ae.arg)
			fmt.Println(ae.prob)
		}
	}
*/
func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
