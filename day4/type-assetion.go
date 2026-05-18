package main

import "fmt"

func main6() {

	// type assertion means:
	// checking whether dynamic value/type
	// stored inside interface
	// is actually of a certain type

	var i interface{} = "hello"

	// direct type assertion

	s := i.(string)

	fmt.Println(s)

	// type assertion can also return:
	// value, ok

	// value -> actual converted value
	// ok    -> whether assertion succeeded

	s, ok := i.(string)

	fmt.Println(s, ok)

	// here assertion fails
	// because interface currently stores string
	// not float64

	f, ok := i.(float64)

	fmt.Println(f, ok)

	// when assertion fails:
	// value becomes zero value of asserted type
	// ok becomes false

	// if we do NOT receive/use ok variable,
	// failed assertion causes runtime panic

	// f = i.(float64) // panic

	// fmt.Println(f)

	//------------------------------------------
	do(21)
	do("hello")
	do(true)
}

func do(i interface{}) {

	// type switch

	// i.(type) can ONLY be used inside type switch

	// it returns/checks the actual dynamic type
	// stored inside interface

	switch v := i.(type) {

	case int:

		// here v becomes int
		fmt.Printf("Twice %v is %v\n", v, v*2)

	case string:

		// here v becomes string
		fmt.Printf("%q is %v bytes long\n", v, len(v))

	default:

		// runs if no matching type found
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
