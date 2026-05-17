package main

import "fmt"

// initialize variable with var and its type at the end

// variables in Go are always initialized with zero values
// variables are never uninitialized/garbage like in some other languages

// zero values:
// int     -> 0
// float   -> 0
// bool    -> false
// string  -> ""

var c, python, com bool

func main2() {
	var x int

	// another way of initializing variable with value
	// explicitly defined type
	var j int = 10

	// inferred type
	var l = 10

	var s string
	var f float32

	// shorter way of initializing variable with value
	// := can only be used inside functions
	i := 0

	fmt.Println(x, i, j, l, c, python, com)

	// %v -> value
	// %q -> string with quotes
	// %T -> type

	fmt.Printf("%v , %v , %T, %q \n", x, f, c, s)
}