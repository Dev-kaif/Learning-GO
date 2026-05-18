package main

import "fmt"

// Generics in Go

// syntax:
//
// func FunctionName[TypeParameter Constraint](input) output

// T -> generic type placeholder
//
// comparable -> constraint
// meaning T must support == and != operators
//
// this allows same function to work
// with multiple data types

// Index returns index of x in slice
// if not found returns -1

func Index[T comparable](s []T, x T) int {

	// s -> slice of type T
	// x -> value of same type T

	for i, v := range s {

		// v and x are both type T

		// because T has comparable constraint,
		// we can safely use == operator

		if v == x {
			return i
		}
	}

	// value not found

	return -1
}

func main() {

	// Index works with int slices

	si := []int{10, 20, 15, -10}

	fmt.Println(Index(si, 15))

	// Index also works with string slices

	ss := []string{"foo", "bar", "baz"}

	fmt.Println(Index(ss, "hello"))
}