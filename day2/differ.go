package main

import "fmt"

func main2() {

	// defer puts the function call into a Stack
	// and executes it at the end of the surrounding function.
	// Deferred calls follow LIFO (Last In First Out).

	fmt.Println("counting")

	for i := range 10 {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	// defer is very useful when panic happens.
	// panic in Go is similar to an exception/runtime crash.
	//
	// Before the function fully exits, Go executes all deferred functions.
	//
	// Example:
	// If you opened a database connection or file,
	// defer ensures it gets closed even if panic occurs.
}