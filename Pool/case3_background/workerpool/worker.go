package workerpool

import (
	"fmt"
)

// Worker Handles all the work
type Worker struct {
	ID        int
	taskChan  chan *Task
	quit      chan bool
	countTask int
}

// NewWorker returns new instance of worker
func NewWorker(channel chan *Task, ID int) *Worker {
	return &Worker{
		ID:        ID,
		taskChan:  channel,
		quit:      make(chan bool),
		countTask: 0,
	}
}

// StartBackground starts the worker in background waiting
func (wr *Worker) StartBackground() {
	fmt.Printf("Starting worker %d\n", wr.ID)
	for {
		select {
		case task := <-wr.taskChan:
			wr.countTask++
			process(wr.ID, task)
		case <-wr.quit:
			return
		}
	}
}

// Stop quits the worker
func (wr *Worker) Stop() {
	fmt.Println("********************************************************")
	fmt.Printf("Closing worker %d - Task processed: %d\n", wr.ID, wr.countTask)
	fmt.Println("********************************************************")
	go func() {
		wr.quit <- true
	}()
}
