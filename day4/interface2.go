package main

import "fmt"

func main5() {

	// empty interface

	// interface{} means: "accept any type"

	// because every type satisfies empty interface
	// as it requires NO methods

	// this way we can store values
	// of any type inside i

	var i interface{}

	i = 4
	i = "hello"

	// similarly Go introduced alias: type any = interface{}
	// so "any" and "interface{}" are same thing
	var j any

	// can store anything
	j = []int{3, 4, 5}

	fmt.Println(i, j)
}