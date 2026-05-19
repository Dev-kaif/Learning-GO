package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe for concurrent use.
//
// Multiple goroutines can access it safely
// because mutex prevents race conditions.
type SafeCounter struct {

	// Mutex
	//
	// Used to LOCK shared data.
	//
	// Only one goroutine can hold
	// lock at a time.
	mu sync.Mutex

	// Shared map data
	v map[string]int
}

// Inc increments value for given key.
func (c *SafeCounter) Inc(key string) {

	/*
		Lock()

		Only ONE goroutine at a time
		can continue past this point.

		Other goroutines trying to lock
		will BLOCK/wait here.

		This prevents multiple goroutines
		from modifying map simultaneously.

		Without mutex:
			race condition can happen.
	*/
	c.mu.Lock()

	/*
		Shared map modification.

		Maps are NOT safe for concurrent writes.

		So mutex protection is needed.
	*/
	c.v[key]++

	/*
		Unlock()

		Releases mutex.

		Now another waiting goroutine
		can acquire lock.
	*/
	c.mu.Unlock()
}

// Value returns current value for key.
func (c *SafeCounter) Value(key string) int {

	// lock before reading shared map
	c.mu.Lock()

	/*
		defer
		
		Unlock() will run automatically
		when function exits.

		Useful because: we won't forget to unlock
	*/
	defer c.mu.Unlock()

	return c.v[key]
}

func main() {

	/*
		Create SafeCounter.

		make(map[string]int)
		creates empty map.
	*/
	c := SafeCounter{
		v: make(map[string]int),
	}

	/*
		Loop itself is NOT concurrent.

		But:

			go c.Inc(...)

		creates NEW goroutine every iteration.

		So:
			1000 goroutines created here.
	*/
	for i := 0; i < 1000; i++ {

		go c.Inc("somekey")
	}

	/*
		Sleep keeps main goroutine alive.

		Otherwise:
			main may finish before
			other goroutines complete.

		This is only for demo.

		Real programs usually use:
			sync.WaitGroup
	*/
	time.Sleep(time.Second)

	//Read final value safely.
	fmt.Println(c.Value("somekey"))
}
