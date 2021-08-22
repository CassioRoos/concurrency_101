package main

import (
	"fmt"
	"time"

	"Concurrency/Pool/case2_worker_pool/workerpool"
)

func main() {
	var allTask []*workerpool.Task
	for i := 0; i < 100; i++ {
		task := workerpool.NewTask(func(data interface{}) error {
			taskID := data.(int)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Task %d processed\n", taskID)
			return nil
		}, i)
		allTask = append(allTask, task)
	}
	pool := workerpool.NewPool(allTask, 5)
	start := time.Now().UTC()
	pool.Run()
	fmt.Println("Took :",time.Since(start))
}
