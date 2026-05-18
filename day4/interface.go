package main

import "fmt"

// interface containing method named Abs
// any type implementing this method satisfies the interface

type Abser interface {
	Abs() int
}

// custom type

type MyInt2 int

// method on custom type

func (i MyInt2) Abs() int {

	if i < 0 {
		return int(-i)
	}

	return int(i)
}

// struct

type Vertex3 struct {
	X, Y int
}

// pointer receiver method on struct

func (v *Vertex3) Abs() int {
	return v.X * v.Y
}

func main4() {

	// nil interface
	var a Abser
	// interface internally stores:
	// 1) dynamic type
	// 2) dynamic value
	
	// currently both are nil


	// instances of custom types

	i := MyInt2(3)

	v := Vertex3{2, 4}

	// MyInt2 has:
	// Abs() int

	// so it satisfies Abser interface

	a = i

	// interface now internally contains:
	// dynamic type  -> MyInt2
	// dynamic value -> 3

	fmt.Println(a.Abs())

	// Vertex3 has pointer receiver method:
	// func (v *Vertex3) Abs() int

	// so ONLY *Vertex3 satisfies interface

	a = &v

	// interface now internally contains:
	// dynamic type  -> *Vertex3
	// dynamic value -> address of v

	fmt.Println(a.Abs())

	// this gives compile error:

	// a = v

	// because Vertex3 value itself
	// does NOT have Abs() method

	// only *Vertex3 has it

	//-------------------------------------------

	// here we directly assign custom type
	// to interface variable

	// variable m is of interface type Abser

	// MyInt2 satisfies Abser interface
	// because it implements:
	// Abs() int

	// interfaces are used to define expected behavior/type

	// they only allow values/types
	// which implement required methods

	// any type implementing required methods
	// can satisfy interface

	var m Abser = MyInt2(4)

	// interface now contains:
	// dynamic type  -> MyInt2
	// dynamic value -> 4

	fmt.Println(m.Abs())
}
