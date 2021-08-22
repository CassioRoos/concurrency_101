package workerpool

import "sync"

// Pool is the worker pool
type Pool struct {
	Task []*Task

	concurrency int
	collector   chan *Task
	wg          sync.WaitGroup
}

// NewPool initializes a new pool with the given tasks and
// at the given concurrency.
func NewPool(tasks []*Task, concurrency int) *Pool{
	return &Pool{
		Task:        tasks,
		concurrency: concurrency,
		collector:   make(chan *Task, 100),
	}
}

// Run runs all work within the pool and blocks until it's finished.
func (p *Pool) Run() {
	for i := 0; i <= p.concurrency; i++ {
		worker := NewWorker(p.collector, i)
		worker.Start(&p.wg)
	}

	for _, task := range p.Task{
		p.collector <- task
	}
	close(p.collector)
	p.wg.Wait()
}