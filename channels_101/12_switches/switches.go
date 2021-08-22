package main

import (
	"fmt"
	"time"
)

func main() {
	for c := make(chan struct{}, 1); true; {
		select {
		case c <- struct{}{}:
			fmt.Print("1")
		case <-c:
			fmt.Print("2")
		}
		time.Sleep(time.Second)
	}
}