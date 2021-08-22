package basic

import (
	"fmt"
	"time"

	"Concurrency/Pool/case1/model"
)

func Work(allData []model.SimpleData) {
	start := time.Now()
	for i, _ := range allData {
		Process(allData[i])
	}
	fmt.Printf("Took ===============> %s\n", time.Since(start))
}

func Process(data model.SimpleData) {
	fmt.Printf("Start processing %d\n", data.ID)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Finish processing %d\n", data.ID)
}
