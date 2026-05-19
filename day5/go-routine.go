package main

import (
	"fmt"
	"time"
)

/*
	Node.js vs Go Concurrency

	Node.js:

	- JavaScript runs mostly on ONE main thread.
	- Uses:
		- Event Loop
		- Callback Queue
		- Async I/O

	When user1 is waiting for something
	(database/API/file/etc): Node DOES NOT block whole server.

	Instead:

		1. Operation is given to runtime/libuv
		2. JS thread becomes free again
		3. Event loop continues handling user2/user3/etc
		4. When result is ready,
		   callback/promise is pushed back into queue

	So Node achieves concurrency mainly using non-blocking async operations

	------------------------------------------------

	Go:

	- Go uses goroutines

	- Goroutines are lightweight execution flows
	  managed by Go runtime itself.

	- Go runtime has a scheduler.

	- Scheduler decides:
		- which goroutine runs
		- when it runs
		- which OS thread runs it
		- when execution switches happen

	- Goroutines are NOT directly:
		1 goroutine = 1 OS thread

	------------------------------------------------

	Traditional Languages:

	In many languages,
	concurrency is often handled using
	OS threads directly.

	When new tasks/users come:
		new threads may be spawned.

	OS threads are expensive because they need:
		- stack memory
		- OS scheduling
		- thread metadata

	Thread stack size is usually much larger
	compared to goroutines.

	Often hundreds of KBs to MBs
	depending on OS/language/runtime.

	So with MANY users/tasks:
		RAM usage can grow heavily.

	------------------------------------------------

	Goroutines: Goroutines are much cheaper.

	They start with very  small stacks
	(roughly few KBs).

	Their stack can dynamically grow/shrink
	when needed.

	This is one reason Go can handle
	thousands or even millions of goroutines.

	------------------------------------------------

	Go runtime multiplexes MANY goroutines
	onto available OS threads efficiently.

	So:

		Many Goroutines
			 ↓
		Go Scheduler
			 ↓
		Fewer/Available OS Threads
			 ↓
		CPU Execution

	------------------------------------------------

	Important:

	When one goroutine blocks/waits:
		- database
		- API
		- file
		- sleep
		- network

	Go scheduler can run OTHER goroutines.

	So whole application does NOT freeze.

	The scheduler maps MANY goroutines onto
	FEWER OS threads efficiently.

	This is called: M:N Scheduling

	------------------------------------------------

	In Go:

	Usually each user/request/task gets
	its own goroutine.

	Example:
		go handleUser(user1)
		go handleUser(user2)

	If user1 waits for:
		- database
		- network
		- file
		- sleep
	etc...

	then ONLY that goroutine pauses/block.

	Meanwhile:
		Go scheduler runs other goroutines.

	So user2 can continue processing
	without freezing whole application.

	------------------------------------------------

	Important:

	Goroutines are NOT exactly tiny OS threads.

	They are:
		lightweight user-space managed tasks

	handled by Go runtime.

	Many goroutines can exist because they are
	much cheaper than OS threads.
*/

/*
	say()

	Simple function that prints a string 5 times.

	time.Sleep() is added so execution becomes slower,
	allowing us to SEE concurrency properly.

	Without Sleep:
	one function may finish too quickly before
	the scheduler switches goroutines.
*/

func say(s string) {

	for i := 0; i < 5; i++ {

		/*
			Pause THIS goroutine for 100 milliseconds.

			Important:
			This does NOT block the whole program.

			Only the current goroutine sleeps.

			Meanwhile:
			other goroutines can continue running.
		*/
		time.Sleep(100 * time.Millisecond)

		// print passed string
		fmt.Println(s)
	}
}

func main1() {

	/*
		go keyword creates a goroutine.

		Goroutine = lightweight execution flow
		managed by Go runtime.

		This does NOT mean:
		"run instantly right now"

		It means:
		"schedule this function to run concurrently"

		So Go runtime puts this goroutine into
		its scheduler system.

		Main goroutine continues immediately
		without waiting.

		task:
			say("world")
	*/
	go say("world")

	/*
		This runs normally inside main goroutine.

		Main goroutine and "world" goroutine
		now run concurrently.

		Both can progress independently.

		Scheduler decides:
		- when each goroutine gets CPU time
		- how execution switches happen

		Output order is NOT guaranteed.
	*/
	say("hello")

	/*
		Important:

		main() itself is also a goroutine
		(main goroutine).

		When main() finishes:
			entire program exits

		Even if other goroutines are still running.
	*/
}