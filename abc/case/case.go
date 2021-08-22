package main

import (
	"fmt"
	"time"
)


func main() {

	mySlice := []string{"A", "B", "C"}
	//for index, value := range mySlice {
	//	go func() {
	//		fmt.Printf("Index: %d\n", index)
	//		fmt.Printf("Value: %s\n", value)
	//	}()
	//}
	//
	//for index, value := range mySlice {
	//	index := index
	//	value := value
	//	go func() {
	//		fmt.Printf("Index: %d\n", index)
	//		fmt.Printf("Value: %s\n", value)
	//	}()
	//}
	for index, value := range mySlice {
		go func(index int, value string) {
			fmt.Printf("Index: %d\n", index)
			fmt.Printf("Value: %s\n", value)
		}(index, value)
	}
	time.Sleep(1*time.Second)
}
