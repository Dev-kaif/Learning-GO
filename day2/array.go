package main

import "fmt"

// Go is a statically typed language

// size is part of an array's type
// [3]string and [2]string are different types

// arrays have fixed size
// size cannot be changed once declared

// due to this, arrays are not used very often directly in real-world Go
// because we usually don't know the exact size of incoming data beforehand

// slices are used more commonly because they are dynamic/flexible

func main5() {
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Zero value of array
	// s := [6]int{} // output : [0 0 0 0 0 0]

	fmt.Println(primes)
}