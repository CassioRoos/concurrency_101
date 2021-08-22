package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Pass send-only channels as arguments

Same as the last example, in the following example, the values of two arguments
of the sumSquares function call are requested concurrently.
Different to the last example, the longTimeRequest function takes a send-only
channel as parameter instead of returning a receive-only channel result.
*/

func longTimeRequest(r chan<- int32) {
	// Simulate a workload.
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func sumSquares(a, b int32) int32 {
	fmt.Printf("a = %d - b = %d\n", a, b)
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// it could be buffered
	ra := make(chan int32)
	go longTimeRequest(ra)
	go longTimeRequest(ra)

	fmt.Println(sumSquares(<-ra, <-ra))
}
