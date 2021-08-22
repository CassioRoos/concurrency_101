package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Use Channels as Futures/Promises
Futures and promises are used in many other popular languages. They are often a
ssociated with requests and responses.

Return receive-only channels as results
In the following example, the values of two arguments of the sumSquares function call
are requested concurrently. Each of the two channel receive operations will block
until a send operation performs on the corresponding channel. It takes about three
seconds instead of six seconds to return the final result.
*/
func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		// Simulate a workload.
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSquares(<-a, <-b))
}
