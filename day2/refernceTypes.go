package main

import "fmt"

// & is used to get the address of a variable
// when I do p = &j , i am assigning the address of j to p
// this is called taking a reference / getting the address

// * is used to point to the value at an address
// when I do *p , i am pointing to the value at the address stored in p
// this is called dereferencing

func main3() {
	i, j := 42, 100

	p := &i // p stores address of i

	fmt.Println("*p =", *p)

	*p = 54 // changes value at the address p is pointing to

	fmt.Println("i =", i)

	p = &j // now p stores address of j

	*p = *p / 10 // change value at address

	fmt.Println("j =", j)
}