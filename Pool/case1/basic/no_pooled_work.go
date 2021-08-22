package basic

import (
	"fmt"
	"sync"
	"time"

	"Concurrency/Pool/case1/model"
)

func NotPooledWork(allData []model.SimpleData) {
	start := time.Now()
	var wg sync.WaitGroup
	workerPoolSize := 100

	dataCh := make(chan model.SimpleData, workerPoolSize)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range dataCh {
			wg.Add(1)
			go func(data model.SimpleData) {
				defer wg.Done()
				Process(data)
			}(data)
		}
	}()

	for i, _ := range allData {
		dataCh <- allData[i]
	}

	close(dataCh)
	wg.Wait()
	fmt.Printf("Took ===============> %s\n", time.Since(start))
}