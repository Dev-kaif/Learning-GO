package main

import "fmt"

func main1() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// slicing using indexes -> [start:end]

	// start index is included
	// end index is excluded

	var s []int = primes[1:4] // output: [3 5 7]

	fmt.Println(s)

	// slices are often called reference-like types
	// because they do not store data themselves
	// they refer to an underlying array

	// when slicing, a new array/copy is NOT created
	// slice points to the same underlying array
	// just with different start/end indexes

	// so changes made through slice
	// also affect the original array

	s[0] = 100

	fmt.Println(s)      // output: [100 5 7]
	fmt.Println(primes) // output: [2 100 5 7 11 13]

	// in Go we can create slices without explicitly defining underlying array

	k := []int{3, 4, 5, 5, 6}

	// difference:
	// array  -> size must be defined
	// [5]int

	// slice -> size not needed
	// []int

	// Go automatically creates the underlying array in background

	fmt.Println(k)

	// we can create slices from slices as well
	// because slices are still pointing to an underlying array

	l := k[1:4]

	fmt.Println(l) // output: [4 5 5]

	// changes in new slice also affect original slice/array

	l[0] = 999

	fmt.Println(l) // output: [999 5 5]
	fmt.Println(k) // output: [3 999 5 5 6]

	// we can create slices of structs as well

	j := []struct {
		x int
		y bool
	}{
		{1, true},
		{2, true},
		{3, false},
	}

	fmt.Println(j)

	// the zero value of reference-like types in Go is nil

	var p []int

	fmt.Printf("%v\n", p)
	fmt.Println(p == nil)

	// make creates slice along with underlying array
	// so the slice will NOT be nil

	// syntax:
	// make(type, length, capacity)

	// type     -> slice type
	// length   -> current usable size of slice
	// capacity -> size of underlying array from slice start point

	a := make([]int, 5) // output: [0 0 0 0 0]

	fmt.Println(a)
	fmt.Println(len(a))   // 5
	fmt.Println(cap(a))   // 5
	fmt.Println(a == nil) // false

	// length can be smaller than capacity
	b := make([]int, 0, 4)
	// length can NEVER be greater than capacity

	fmt.Println(b)
	fmt.Println(len(b)) // 0
	fmt.Println(cap(b)) // 4

	// here capacity is still 4
	// because this slice starts from index 0
	// so it can still access till the end of underlying array

	c := b[0:2]

	fmt.Println(c)
	fmt.Println(cap(c)) // 4

	// here capacity becomes smaller
	// because slice now starts from index 2
	d := b[2:4]

	// capacity is counted from the slice starting index
	// till the end of the underlying array

	fmt.Println(d)
	fmt.Println(cap(d)) // 2

	// --------------- 2D slices -----------------

	// slice of slices
	// each element stores another slice

	board := [][]string{
		[]string{"__", "__", "__"},
		[]string{"__", "__", "__"},
		[]string{"__", "__", "__"},
	}

	board[0][0] = "X"
	board[2][2] = "X"
	board[1][1] = "X"

	fmt.Println(board)
}
