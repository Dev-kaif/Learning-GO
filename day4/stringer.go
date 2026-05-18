package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// Stringer

// Go internally has an interface: fmt.Stringer

// it contains method: String() string

// if a type satisfies this interface,
// fmt.Println and other fmt functions
// use this String() method for printing

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age) + "\nRandom words lol\n"
}

func main7() {

	a := Person{"Arthur Dent", 42}

	z := Person{"Zaphod Beeblebrox", 9001}

	// since Person satisfies fmt.Stringer,
	// Println uses String() method internally

	fmt.Println(a, z)

	// if type does NOT satisfy Stringer,
	// Go uses default formatting behavior
}
