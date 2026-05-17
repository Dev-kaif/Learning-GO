package main

import (
	"fmt"
	"math"
)

// function taking another function as input

// func(float64, float64) float64
// is the function type/signature accepted

// meaning:
// takes 2 float64 values
// returns 1 float64 value

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {

	// anonymous function
	// function without a name

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	// passing function to another function

	fmt.Println(compute(hypot))

	// func math.Pow(x float64, y float64) float64 
	fmt.Println(compute(math.Pow))
}
