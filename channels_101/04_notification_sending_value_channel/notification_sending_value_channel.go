package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"sort"
	"time"
)
/*
1-to-1 notification by sending a value to a channel
If there are no values to be received from a channel, then the next receive operation
on the channel will block until another goroutine sends a value to the channel.
So we can send a value to a channel to notify another goroutine which is waiting
to receive a value from the same channel.

In the following example, the channel done is used as a signal channel to do
notifications.
*/
func main() {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // can be buffered or not

	// The sorting goroutine
	go func() {
		start := time.Now()
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		fmt.Println("goroutine took:", time.Since(start))
		// Notify sorting is done.
		//done <- struct{}{}
		// 'Close' will do the work, and it is more commonly used
		close(done)
	}()

	// do some more stuff
	<- done
	fmt.Println(len(values))
	// uncomment the line bellow in order to print everything
	//fmt.Println(values, values[len(values)-1])
}
