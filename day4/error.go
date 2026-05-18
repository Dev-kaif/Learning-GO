package main

import (
	"fmt"
	"time"
)

/*
Go already has a built-in error interface

type error interface {
	Error() string
}

Any type/struct having:
Error() string

automatically satisfies/implements
the error interface

When errors are printed/logged using:
fmt.Println(err)

Go internally calls:
err.Error()

to get string representation of error

This lets us create our own custom error types
with extra information
*/

// custom error struct
type MyError struct {
	When time.Time
	What string
}

// custom implementation of Error() method

// because MyError has:
// Error() string

// it now satisfies Go's error interface
// and can be used anywhere normal errors are used
func (e *MyError) Error() string {

	return fmt.Sprintf(
		"at %v, %s",
		e.When,
		e.What,
	)
}

// function returning error type
func run() error {

	// returning custom error
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main8() {

	// err stores returned error

	// if err is NOT nil:
	// error occurred

	// if err == nil:
	// program continues normal/success flow

	if err := run(); err != nil {

		// fmt.Println internally calls:
		// err.Error()

		fmt.Println(err)

	} else {

		fmt.Println("success")
	}
}
