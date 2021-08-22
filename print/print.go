package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(swap(i, i+1))
	}
}

func swap(a, b int) (int, int) {
	return b, a
}
