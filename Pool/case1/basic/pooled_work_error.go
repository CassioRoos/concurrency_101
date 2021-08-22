package basic

import (
	"fmt"
	"sync"
	"time"

	"Concurrency/Pool/case1/model"
)

func PooledWorkError(allData []model.SimpleData) {
	start := time.Now().UTC()
	var wg sync.WaitGroup
	workerPoolSize := 100

	dataCh := make(chan model.SimpleData, workerPoolSize)
	errorsCh := make(chan error, len(allData))

	for i := 0; i < workerPoolSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range dataCh {
				process(data, errorsCh)
			}
		}()
	}

	for _, data := range allData {
		dataCh <- data
	}
	close(dataCh)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case err := <-errorsCh:
				fmt.Println("finished with error:", err.Error())
			case <-time.After(time.Second * 1):
				fmt.Println("Timeout: errors finished")
				return
			}
		}
	}()
	defer close(errorsCh)
	wg.Wait()
	fmt.Printf("Took ===============> %s\n", time.Since(start))
}

func process(data model.SimpleData, errors chan<- error) {
	fmt.Printf("Start processing %d\n", data.ID)
	time.Sleep(100 * time.Millisecond)
	if data.ID%11 == 0 {
		errors <- fmt.Errorf("error on job %v", data.ID)
	} else {
		fmt.Printf("Finish processing %d\n", data.ID)
	}
}
