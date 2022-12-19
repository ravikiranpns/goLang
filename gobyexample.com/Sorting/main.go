/*

https://gobyexample.com/sorting
Sorting
- Go's sort package implements sorting for builtins and user-
defined types.

- Sort methods are specific to the builtin type;
- Ex - strings.
- Note - the sorting is in-place, so it changes the given slice
and doesn't return a new one.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

- Ex - ints



- use sort to check if a slice is already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)



*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}

	s := sort.StringsAreSorted(strs)
	fmt.Println("Strings Sorted: ", s)

	sort.Strings(strs)
	fmt.Println("Strings: ", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:  ", ints)

	s = sort.IntsAreSorted(ints)
	fmt.Println("Ints Sorted: ", s)

	s = sort.StringsAreSorted(strs)
	fmt.Println("Strings Sorted: ", s)
}
