// https://gobyexample.com/generics
// ONS - 1822600.06.10.22
package main

/* Starting with version 1.18, Go added Support for
generics, also known as type parameters.

Ex.Generic functions
- MapKeys takes a map of any type and returns a slice of its keys.
- Has two type parameters -K and V
	func MapKeys[K comparable, V any](m map[K]V) []K {
		r := make([]K, 0, len(m))
		for k := range m {
			r = append(r, k)
		}
		return r
	}

- K has the comparable constraint(we can compare values of this type
with == and != ops). Required for map keys in Go.

- V has the any constraint(not restricted in any way, alias for interface{})

- An example(generic type), List is a SLL with values of any type
	type List[T any] struct {
		head, tail *element[T]
	}

	type element[T any] struct {
		next *element[T]
		val T
	}

- Define mehtods on generic types, as regular types, but we have to keep
ty type parameters in place. The type is List[T], not List.
	func (lst *List[T]) Push(v T) {
		if lst.tail == nil {
			lst.head = &element[T]{val: v}
			lst.tail = lst.head
		} else {
			lst.tail.next = &element[T]{val: v}
			lst.tail = lst.tail.next
		}
	}

	func (lst *List[T]) GetAll() []T {
		var elems []T
		for e := lst.head; e != nil; e = e.next {
			elems = append(elems, e.val)
		}
		return elems
	}

- Invoking generic functions, often rely on type inference.
NOTE - Don't have to specify the types for K and V when calling MapKeys
the compiler infers them automatically.
	func main() {
		var m = map[int]string{1: "2", 2 : "4", 4: "8"}

		fmt.Println("keys: ", MapKeys(m))

		_= MapKeys[int, string](m)

		lst := List[int]{}
		lst.Push(10)
		lst.Push(13)
		lst.Push(23)
		fmt.Println("list: ", lst.GetAll())
	}
*/

// func MapKeys[K comparable, V any](m map[K]V) []K {
// 	r := make([]K, 0, len(m))
// 	for k := range m {
// 		r = append(r, k)
// 	}
// 	return r
// }

// type List[T any] struct {
// 	head, tail *element[T]
// }

// type element[T any] struct {
// 	next *element[T]
// 	val  T
// }

// func (lst *List[T]) Push(v T) {
// 	if lst.tail == nil {
// 		lst.head = &element[T]{val: v}
// 		lst.tail = lst.head
// 	} else {
// 		lst.tail.next = &element[T]{val: v}
// 		lst.tail = lst.tail.next
// 	}
// }

// func (lst *List[T]) GetAll() []T {
// 	var elems []T
// 	for e := lst.head; e != nil; e = e.next {
// 		elems = append(elems, e.val)
// 	}
// 	return elems
// }

func main() {
	// var m = map[int]string{1: "2", 2: "4", 4: "8"}

	// fmt.Println("keys: ", MapKeys(m))

	// _ = MapKeys[int, string](m)

	// lst := List[int]{}
	// lst.Push(10)
	// lst.Push(13)
	// lst.Push(23)
	// fmt.Println("list: ", lst.GetAll())
}
