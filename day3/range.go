package main

import "fmt"

var paw = []int{2, 3, 4, 5, 6, 7, 8}

func main3() {

	// range returns: index (i) and value (v)

	// value v is a COPY of the original value and i is the index of that value
	// so operations on v do NOT affect actual slice values

	for i, v := range paw {

		fmt.Println("i:", i, " v:", v)

		// this only changes copied value
		// original slice remains unchanged from this operation

		v = v * 100

		// this affects original slice
		// because we are modifying actual index

		paw[i] = paw[i] * 10
	}

	fmt.Println(paw)
}
