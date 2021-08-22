package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
This is the enhancement of the using-only-one-channel variant in the last example.

Sometimes, a piece of data can be received from several sources to avoid high
latencies. For a lot of factors, the response durations of these sources may
vary much. Even for a specified source, its response durations are also not constant.
To make the response duration as short as possible, we can send a request to every
source in a separated goroutine. Only the first response will be used, other slower
ones will be discarded.

Note, if there are N sources, the capacity of the communication channel must
be at least N-1, to avoid the goroutines corresponding the discarded responses
being blocked for ever.
*/

var names = []string{"athena", "zeus", "hercules", "hades", "hephaestus"}

func source(c chan<- string, pos int) {
	rb := rand.Intn(3) + 1
	// Sleep 1s/2s/3s.
	fmt.Println(
		"going to sleep for",
		rb, "seconds",
		"my position on the array is",
		pos,
	)
	time.Sleep(time.Duration(rb) * time.Second)
	c <- names[pos]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	startTime := time.Now()

	//c must be a buffered channel
	c := make(chan string, 5)
	for i := 0; i < 5; i++ {
		go source(c, i)
	}
	rnd := <-c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}
