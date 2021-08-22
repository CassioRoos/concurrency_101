package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
N-to-1 and 1-to-N notifications

By extending the above two use cases a little, it is easy to do N-to-1 and 1-to-N
notifications.

In fact, the ways to do 1-to-N and N-to-1 notifications introduced in this
sub-section are not used commonly in practice. In practice, we often use
sync.WaitGroup to do N-to-1 notifications, and we do 1-to-N notifications by
close channels. Please read the next sub-section for details.
*/

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // block here and wait a notification
	log.Println("Worker", id, "starts.")
	// Simulate a workload
	r := rand.Intn(3) + 2
	fmt.Println("this will take",r,"seconds")
	time.Sleep(time.Duration(r) *time.Second)
	log.Print("Worker", id, " job done.")
	// Notify the main goroutine (N-to-1),
	done <- T{}
}

func main() {
	log.SetFlags(0)
	rand.Seed(time.Now().UnixNano())
	ready, done := make(chan T), make(chan T)

	for i := 0; i < 3; i++ {
		go worker(i, ready, done)
	}

	// Simulate an initialization phase.
	time.Sleep(time.Second * 3 / 2)


	// 1-to-N notifications.

	/*
	The way to do 1-to-N notifications shown in the last sub-section is seldom used in
	practice, for there is a better way. By making using of the feature that infinite
	values can be received from a closed channel, we can close a channel to broadcast
	notifications.

	By the example in the last sub-section, we can replace the three channel send
	operations ready <- struct{}{} in the last example with one channel close
	operation close(ready) to do an 1-to-N notifications.
	*/
	close(ready)
	//for i := 0; i < 3; i++ {
	//	ready <- T{}
	//}


	// Being N-to-1 notified.
	for i := 0; i < 3; i++ {
		<-done
	}
}
