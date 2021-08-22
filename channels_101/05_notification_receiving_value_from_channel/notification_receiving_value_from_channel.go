package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	// The capacity of the signal channel can
	// also be one. If this is true, then a
	// value must be sent to the channel before
	// creating the following goroutine.
	go func() {
		fmt.Println("This will take some time")
		// Simulate a workload
		time.Sleep(3 * time.Second)

		// Receive a value from the done
		// channel, to unblock the second
		// send in main goroutine.
		<- done
	}()

	// Blocked here, wait for the notification
	done <- struct{}{}
	// By sending a value, we will block the channel until the value is used
	fmt.Println("All done!")
}
