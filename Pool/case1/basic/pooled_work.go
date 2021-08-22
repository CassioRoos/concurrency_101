package basic

import (
	"fmt"
	"sync"
	"time"

	"Concurrency/Pool/case1/model"
)

func PooledWork(allData []model.SimpleData){
	start := time.Now().UTC()
	var wg sync.WaitGroup
	workerPoolSize := 100

	dataCh := make(chan model.SimpleData, workerPoolSize)

	for i := 0; i < workerPoolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range dataCh {
				Process(data)
			}
		}()
	}

	for i, _ := range allData{
		dataCh <- allData[i]
	}

	close(dataCh)
	wg.Wait()
	fmt.Printf("Took ===============> %s\n", time.Since(start))
}
