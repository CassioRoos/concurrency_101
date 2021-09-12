package main

import "fmt"

func main() {
	c := make(chan int, 2) // a buffered channel
	c <- 3
	c <- 5
	//c <- 10 // panic!
	close(c)
	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c                // Reads from the channel, unshifts it, gets the OK because if there is nothing inside, as an int it would still be a number 0
	fmt.Println(x, ok)          // 3 true
	fmt.Println(len(c), cap(c)) // 1 2
	x, ok = <-c
	fmt.Println(x, ok)          // 5 true
	fmt.Println(len(c), cap(c)) // 0 2
	x, ok = <-c                 // channel of type int so it is a 0 received but, ok = false
	fmt.Println(x, ok)          // 0 false
	x, ok = <-c
	fmt.Println(x, ok)          // 0 false
	fmt.Println(len(c), cap(c)) // 0 2
	close(c)                    // panic! cannot close a closed channel
	// The send will also panic if the above
	// close call is removed.
	c <- 7 // panic! cannot send a value to a closed channel
}
