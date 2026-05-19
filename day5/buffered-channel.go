package main

import "fmt"

/*
	------------------------------------------------

	BUFFERED vs UNBUFFERED CHANNELS

	1. Unbuffered Channel

		make(chan int)

		- Has NO internal storage/buffer.
		- Sender and receiver must synchronize directly.

		Send:

			c <- value

		blocks until receiver takes value.

		Receive:

			<-c

		blocks until sender sends value.

		So unbuffered channels act like:
			direct hand-to-hand transfer.

		------------------------------------------------

		2. Buffered Channel

			make(chan int, size)

		Example:

			make(chan int, 3)

		- Has internal queue/buffer.
		- Buffer follows FIFO order.
		- Values can wait inside channel.

		Sender blocks ONLY when:
			buffer becomes full.

		Receiver blocks ONLY when:
			buffer becomes empty.

		Buffered channels are useful when:
			sender and receiver do not need
			immediate synchronization.

	------------------------------------------------
*/

/*
	Important:

	main() itself runs inside the
	main goroutine.

	If main goroutine finishes execution,
	the ENTIRE program exits immediately.

	This means:

		If other goroutines are still running,
		they are terminated automatically.

	So main goroutine often needs to wait
	for other goroutines to complete.

	------------------------------------------------

	Channels help with this synchronization.

	Example:

		x := <-c

	Main goroutine blocks/waits here until
	some goroutine sends data into channel.

	So channel communication can naturally
	prevent main() from finishing too early.

	------------------------------------------------

	In this example:

		go sum(...)

	starts worker goroutines.

	Then:

		x, y := <-c, <-c

	forces main goroutine to wait until
	BOTH goroutines send their results.

	Only after receiving both values:

		main() continues execution.

	This is why output is not lost.
*/

func main3() {

	/*
		Buffered channel with capacity 2.

		Can store 2 values internally
		before blocking sender.

		Follows FIFO order.
	*/
	ch := make(chan int, 2)

	// buffer: [1]
	ch <- 1

	// buffer: [1, 2]
	ch <- 2

	/*
		If we do:

			ch <- 3

		it will block because
		buffer is full.
	*/

	/*
		Receive removes/consumes value
		from channel.

		FIFO:
			first inserted comes out first.
	*/
	fmt.Println(<-ch) // 1

	// buffer now: [2]
	fmt.Println(<-ch) // 2

	// buffer now: []
}
