package main

import (
	"fmt"
	"time"
)

func main5() {

	// store starting time
	start := time.Now()

	// time.Tick() RETURNS a channel

	// every 100ms Go runtime sends
	// current time into this channel
	tick := time.Tick(100 * time.Millisecond)

	// time.After() also returns channel

	// after 500ms it sends ONE value
	boom := time.After(500 * time.Millisecond)

	// helper function for elapsed time
	elapsed := func() time.Duration {
		return time.Since(start).Round(time.Millisecond)
	}

	for {

		// select works like switch
		// but for channels
		select {

		// receive from tick channel
		// equivalent: t := <-tick
		case <-tick:
			fmt.Printf("[%6s] tick.\n", elapsed())

		// receive from boom channel
		// runs once after 500ms
		case <-boom:
			fmt.Printf("[%6s] BOOM!\n", elapsed())
			return

		// default runs when no channel is ready yet
		default:
			fmt.Printf("[%6s]     .\n", elapsed())
			time.Sleep(50 * time.Millisecond)
		}
	}
}
