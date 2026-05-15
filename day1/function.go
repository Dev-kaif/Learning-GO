package main

import "fmt"

// defined the input type and return type
func add(a int, b int) int {
	return a + b
}

// multi return
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	// already insiliazed when declared type of return with varible names x =0 , y = 0
	x = sum / 2
	y = sum / 3
	return
}

func split2(sum int) (int, int) {
	// didn't insiliaze when declared type of return
	x := sum / 2
	y := sum / 3
	return x, y
}

func main1() {

	fmt.Println(add(2, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	c, d := split(67)
	fmt.Println(c, d)

}
