package main

import "fmt"

func main2() {

	// nil slice
	// len = 0
	// cap = 0
	var s []int

	fmt.Println(cap(s))

	// append adds new elements to slice

	// append returns a NEW updated slice
	// because underlying array may be replaced internally

	// if capacity becomes full,
	// Go creates a new bigger underlying array,
	// copies old values into it,
	// and returned slice points to this new array

	// usually capacity gets doubled (2x) till around capacity 256,

	// example:
	// 1 -> 2 -> 4 -> 8 -> 16

	// after that Go grows capacity more slowly roughly around 1.25x

	// this helps balance memory usage and reallocation performance

	// so we must reassign it back

	s = append(s, 4)

	fmt.Println(s, cap(s))

	// append can take multiple elements as input

	s = append(s, 4, 5, 6, 7, 8)

	fmt.Println(s, cap(s))
}
