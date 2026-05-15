package main

import "fmt"

// initialise varible with var and its type at the end
// varible in Go is always initialise in zero state
// varible is never initlaised in null state
var c, python, com bool

func main2() {
	var x int

	// another way of initializing the variable with value
	// defined type
	var j int = 10

	// intrupted type
	var l = 10

	var s string
	var f float32

	// shoter way of initializing the variable with value
	i := 0

	fmt.Println(x, i, j, l, c, python, com)

	// %v is value , %q prints string with qoute , %T shows type
	fmt.Printf("%v , %v , %T, %q \n", x, f, c, s)
}
