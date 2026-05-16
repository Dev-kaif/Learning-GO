package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main1() {
	// Normal if else conditional
	v := rand.Intn(10)
	if v <= 5 {
		fmt.Printf("Number is %v\n", v)
	} else {
		fmt.Printf("Number is above '5' = %v\n", v)
	}

	// same but varible is initalised before condition in condition
	if j := rand.Intn(10); j <= 5 {
		fmt.Printf("Number is %v\n", j)
	} else {
		fmt.Printf("Number is above '5' = %v\n", j)
	}

	// switch case
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
