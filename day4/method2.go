package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// pointer receiver method
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// normal function taking pointer argument
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main3() {

	v := Vertex{3, 4}

	// Go automatically converts:
	// v.Scale(2) into (&v).Scale(2)

	// when calling pointer receiver methods
	v.Scale(2)

	// but for normal functions,
	// we must explicitly pass pointer
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}

	// similarly Go automatically converts:
	// p.Scale(3)

	// into:
	// (*p).Scale(3)

	// this is convenience provided by Go for method receivers
	p.Scale(3)

	// normal function already expects pointer
	// so this works directly
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

/*
When to use value receiver vs pointer receiver

1) use pointer receiver when struct size is large
   and you want to avoid copying struct every method call

2) use pointer receiver when you want to mutate/modify original values

3) value receivers are good for:
   read-only methods
   small structs

4) there is a common Go convention:
   if some methods use pointer receivers,
   usually all methods should use pointer receivers

   avoid mixing value + pointer receivers unnecessarily
   because it can create subtle/confusing behavior
*/
