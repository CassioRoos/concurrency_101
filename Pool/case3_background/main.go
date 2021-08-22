package main

import (
	"math/rand"
	"os"
	"os/signal"
	"time"

	"Concurrency/Pool/case3_background/workerpool"
)

func main() {
	var allTask []*workerpool.Task
	// the idea here is just to put some weight on the tasks, we have 5 WORKERS
	for i := 0; i < 10000; i++ {
		task := workerpool.NewTask(func(data interface{}) error {
			//taskID := data.(int)
			time.Sleep(100 * time.Millisecond)
			//fmt.Printf("Task %d processed\n", taskID)
			return nil
		}, i)
		allTask = append(allTask, task)
	}

	// Creates the pool with that task wight on in - 1000 tasks and 5 workers
	pool := workerpool.NewPool(allTask, 5)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)
	go func() {
		// WAIT until the signal comes. This is blocking, then will wait until something occurs
		<-sigChan
		pool.Stop()
	}()

	// runs a go func -- non blocking
	go func() {
		// infinite loop
		for {
			taskID := rand.Intn(100) + 20

			if taskID%99 == 0 {
				pool.Stop()
			}
			// it will create a task every second
			time.Sleep(time.Duration(rand.Intn(1)) * time.Second)
			task := workerpool.NewTask(func(data interface{}) error {
				time.Sleep(100 * time.Millisecond)
				return nil
			}, taskID)

			// adds the task to the pool
			pool.AddTask(task)
		}
	}()

	// As the go func will no block, this peace of code will run right after
	pool.RunBackground()
}
