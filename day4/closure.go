package main

import "fmt"

// function that returns another function
// returned function takes int as input and returns int

func adder() func(int) int {

	// function scoped variable

	sum := 0

	return func(x int) int {

		// due to lexical scope,
		// inner function can access variables
		// from outer function

		sum += x

		return sum

		// due to closure property,
		// even after adder() finishes execution,
		// returned function still remembers/accesses sum

		// sum is kept alive in memory
		// because returned function is still using it
	}
}

func main1() {

	add1 := adder()

	// adder() runs once

	// sum variable becomes part of closure environment

	// even after adder() finishes,
	// sum is NOT destroyed
	// because returned function still references it

	// this persistent remembered state
	// is called a closure

	// add 10 to sum

	fmt.Println(add1(10))

	// current sum = 10
	// adds 103

	fmt.Println(add1(103))

	// current sum = 113
	// adds 7

	fmt.Println(add1(7))
}
