package workerpool

import "fmt"

type Task struct {
	Err  error
	Data interface{}
	f    func(interface{}) error
}

func NewTask(f func(interface{}) error, data interface{}) *Task {
	return &Task{f: f, Data: data}
}

func process(workerID int, task *Task) {
	fmt.Printf("Worker %d process task %v\n", workerID, task.Data)
	task.Err = task.f(task.Data)
}
