package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // an unbuffered channel
	go func(ch chan<- int, x int) {

		// <-ch    // fails to compile
		// Send the value and block until the result is received.
		ch <- x * x // 9 is sent
	}(c, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		// Block until 9 is received.
		n := <-ch
		fmt.Println(n) // 9
		time.Sleep(time.Second * 10)
		// ch <- 123   // fails to compile
		done <- struct{}{}
	}(c)
	// Block here until a value is received by
	// the channel "done".
	<-done
	fmt.Println("bye")
}
