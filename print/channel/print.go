package main

import "fmt"

func main() {
	intChan := make(chan int)
	go gen(intChan)
	for i := range intChan {
		go fmt.Println(swap(i, i+1))
	}
}

func gen(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func swap(a, b int) (int, int) {
	return b, a
}
