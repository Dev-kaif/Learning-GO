package main

import "fmt"

func main4() {

	// for loop
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// while loop made with for loop
	sum2 := 1

	for sum2 < 5 {
		sum2 += sum2
	}

	fmt.Println(sum2)

	// infinite loop
	sum3 := 1
	for {
		sum3 += sum3
	}
}
