package main

import "fmt"

func main() {
	var arr [3]string

	arr[0] = "Om"
	arr[1] = "Shree"
	arr[2] = "Ram"

	fmt.Println("Elements of Arr")
	fmt.Println("1 : ", arr[0])
	fmt.Println("2 : ", arr[1])
	fmt.Println("3 : ", arr[2])
	fmt.Println("Length of arr : ", len(arr))

	// Shorthand declaration
	// arr_name := [len]Type{item1, item2, item3, ...itemN}
	shortArr := [4]string{"go", "goLang", "vscode"}
	fmt.Println("Elements of shortArr: ")

	for i := 0; i < 3; i++ {
		fmt.Println(shortArr[i])
	}

	// Multi-Dimensional Array
	// arr_name[len1][len2]..[lenN]Type
	marr := [3][3]int{{1, 2, 3},
		{5, 6, 7},
		{7, 7, 7},
	}

	fmt.Println("Elements of Multi-Dim Arr:", marr)
	for p := 0; p < 3; p++ {
		for q := 0; q < 3; q++ {
			fmt.Println(marr[p][q])
		}
	}

	ellipsisArr := [...]string{"JaiShreeRam", "JaiShreeKrishna", "JaiShreeGovinda", "JaiShreeBadrinath"}
	fmt.Println("Elements of Ellipsis Arr : ", ellipsisArr)
	for i, item := range ellipsisArr {
		fmt.Println(i, item)
	}
}
