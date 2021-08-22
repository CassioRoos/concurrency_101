package workerpool

import (
	"fmt"
	"time"
)

// Pool is the worker pool
type Pool struct {
	Tasks   []*Task
	Workers []*Worker

	concurrency   int
	collector     chan *Task
	runBackground chan bool
}

// AddTask adds a task to the pool
func (p *Pool) AddTask(task *Task) {
	p.collector <- task
}

// NewPool initializes a new pool with the given tasks and
// at the given concurrency.
func NewPool(tasks []*Task, concurrency int) *Pool {
	return &Pool{
		Tasks:       tasks,
		concurrency: concurrency,
		collector:   make(chan *Task, 100),
	}
}

// RunBackground runs the pool in background
func (p *Pool) RunBackground() {

	go func() {
		for {
			fmt.Println("⌛ Waiting for tasks to come in ...")
			time.Sleep(1 * time.Second)
		}
	}()

	// create the the worker and start them in the background
	for i := 0; i < p.concurrency; i++ {
		worker := NewWorker(p.collector, i)
		p.Workers = append(p.Workers, worker)
		go worker.StartBackground()
	}

	// add all existing task to chan
	for i := range p.Tasks {
		p.collector <- p.Tasks[i]
	}

	// creates the chan
	p.runBackground = make(chan bool)

	// set the work in background
	// this is blocking
	<-p.runBackground
}

// Stop stops background workers
func (p *Pool) Stop() {
	for i := range p.Workers {
		p.Workers[i].Stop()
	}
	p.runBackground <- true
}
