package main

import "fmt"

/*
	CHANNELS IN GO
	------------------------------------------------

	Why channels exist?

	In many traditional languages:

		Threads usually communicate by
		sharing SAME memory/state.

	Example:
		multiple threads modifying same variable.

	This can cause:
		- race conditions
		- synchronization problems
		- locks/mutex complexity

	------------------------------------------------
	Go's Philosophy:

	Instead of:

		"communicate by sharing memory"

	Go prefers:

		"share memory by communicating"

	Meaning:

	Goroutines communicate using CHANNELS.

	Channels are communication pipes
	between goroutines.

	Data flows THROUGH channels.

	------------------------------------------------

	Important Properties of Channels

	1. Channels are typed

		chan int
			-> only int values allowed

	------------------------------------------------

	2. Channels are FIFO

		FIFO = First In First Out

		If values enter channel in order:

			10
			20
			30

		Then receives happen in same order.

	------------------------------------------------

	3. Receiving consumes/removes value

		<-c

		takes value OUT of channel.

		After receive:
			value is gone from channel.

		So channel is NOT permanent storage
		like arrays/slices/database.

	------------------------------------------------

	4. Channel does NOT track sender identity

		If multiple goroutines send into same channel:

			go A()
			go B()

		then receiver only knows:
			which value arrived

		NOT:
			which goroutine sent it

		Go does NOT automatically provide:
			- goroutine id
			- sender identity
			- sender metadata

		If identity is needed:
			we must explicitly send metadata.

		Example:
			type Result struct {
				from string
				value int
			}

	------------------------------------------------

	5. Output order between goroutines
	   is NOT guaranteed

		Because goroutines run concurrently.

		Scheduler decides:
			- which goroutine runs first
			- when it runs
			- when switching happens

		So channel receive order depends on:
			which goroutine sends first.

	------------------------------------------------

	6. Channels synchronize goroutines

		Unbuffered channel send:

			c <- value

		waits until some goroutine receives it.

		Receive:

			<-c

		waits until some goroutine sends value.

		So channels help goroutines
		coordinate execution safely.
*/

func sum(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v
	}

	/*
		Send sum INTO channel.

		<- direction matters.

		c <- sum

		Means:
			send value INTO channel

		Important:

		This send blocks CURRENT goroutine
		until some receiver takes the value.

		So sender waits here until:

			<-c

		happens somewhere else.

		This creates synchronization
		between goroutines.
	*/
	c <- sum
}

func main2() {

	s := []int{7, 2, 8, -9, 4, 0}

	// Create unbuffered channel
	// that transfers int values.
	c := make(chan int)

	// Create buffered channel
	// c2 := make(chan int, 10)

	l := len(s) / 2

	go sum(s[:l], c) // [7 2 8]

	go sum(s[l:], c) //[ -9 4 0]

	/*
		Receive values from channel.

		<-c
			receive value FROM channel

		Important:

		Order is NOT guaranteed.

		We do NOT know:
			which goroutine finishes first.

		So:

			x may become:
				17
			or:
				-5

		depending on scheduler timing.

		------------------------------------------------

		Also important:

		After value is received:
			it is REMOVED from channel.

		So:

			x := <-c

		consumes one value.

		That same value cannot be received again
		unless explicitly sent again.
	*/

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}

