package main

import "fmt"

type Vertex struct {
	X int
	Y int

	// another shorter way:
	// X, Y int
}

var (
	v1 = Vertex{2, 3}

	// named fields
	// Y is implicitly 0
	v2 = Vertex{X: 3}

	// both X and Y implicitly become 0
	v3 = Vertex{}

	// pointer to struct
	q = &Vertex{2, 3}
)

func main4() {
	v := Vertex{2, 4}

	fmt.Println(v)

	p := &v // p stores address of v

	// normally to access value from pointer we dereference:
	// (*p).X = 100

	// but Go gives a shortcut for struct pointers:
	// p.X automatically becomes (*p).X internally

	p.X = 100

	// printing pointer prints memory address
	fmt.Println(p)

	// to print actual value:
	// fmt.Println(*p)

	// ---------------- Zero value of struct ----------------

	var z Vertex

	// all fields get their own zero values automatically
	// int -> 0
	// string -> ""
	// bool -> false

	fmt.Println(z) // output: {0 0}
}
