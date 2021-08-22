package main

import (
	"fmt"

	"Concurrency/Pool/case1/basic"
	"Concurrency/Pool/case1/model"
)

func main() {
	// Prepare the data
	var allData []model.SimpleData
	for i := 0; i < 1000; i++ {
		data := model.SimpleData{ ID: i }
		allData = append(allData, data)
	}
	fmt.Printf("Start processing all work \n")

	// Process
	//basic.Work(allData)
	//basic.PooledWork(allData)
	basic.PooledWorkError(allData)
}