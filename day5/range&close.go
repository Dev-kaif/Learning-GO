package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {

	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}

	/*
		close(c)

		Closes the channel.

		Meaning:
			"No more data will be sent
			 through this channel."

		Important:
		- Existing buffered values can still
		  be received.
		- range loop stops automatically
		  once:
				1. channel is closed
				2. all remaining values are consumed

		- Usually sender closes channel.
		
		- Sending into closed channel
		  causes panic.
	*/
	close(c)
}

func main4() {

	/*
		Buffered channel with capacity 10.

		Can temporarily store 10 values.

		Follows FIFO order.
	*/
	c := make(chan int, 10)

	/*
		cap(c)

		Returns channel capacity.

		Here:
			cap(c) = 10
	*/
	go fibonacci(cap(c), c)

	/*
		range keeps receiving values
		from channel until:

			1. channel is closed
			2. all values are consumed

		Without close(c):
			this loop would wait forever.
	*/
	for i := range c {
		fmt.Println(i)
	}
}
