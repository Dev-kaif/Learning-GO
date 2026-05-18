package main

import "fmt"

// Struct type Vertex
type Vertex1 struct {
	X, Y int
}

// here we are defining a method attached to struct

// to do that first we specify
// which struct this method belongs to

// (v Vertex)
// v -> receiver variable name
// Vertex -> receiver type

// after that we define method normally

// now with receiver variable v,
// we can access struct fields

// here a COPY of original struct is passed
// so original struct cannot be modified

// value receiver
func (v Vertex1) addVertex() int {
	return v.X + v.Y
}

// custom type based on int
type myInt int

// methods can also be attached to custom types
func (v myInt) square() int {
	x := int(v * v)
	return x
}

// passing pointer lets method access actual struct
// instead of copy

// so we can modify original struct values
// pointer receiver
func (v *Vertex1) Scale(x int) {
	v.X = v.X * x
	v.Y = v.Y * x
}

func main2() {

	// initialized the struct

	v := Vertex1{3, 4}

	// creating value of custom type
	s := myInt(3)

	// calling methods
	fmt.Println(v.addVertex(), s.square())

	// pointer vertex
	j := Vertex1{2, 4}

	j.Scale(3)

	fmt.Println(j)
}
