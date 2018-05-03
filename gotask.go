package gotask

import (
	"sync"
	"time"

	log "github.com/lytics/logrus"
)

/*
task func must be is type
func(interface{}, *sync.WaitGroup, chan struct{})
*/

// TaskConfig ....
type TaskConfig struct {
	Handle    func(interface{})
	WorkerNum int
	TimeOut   time.Duration
}

// Task ....
type Task struct {
	Operator     func(interface{}) error
	Args         []interface{}
	WorkerChanel chan struct{}
	Wg           *sync.WaitGroup
	TimeOut      time.Duration
}

// NewTask ....
func (tc *TaskConfig) NewTask(f func(interface{}) error) *Task {
	return &Task{
		Wg:           &sync.WaitGroup{},
		Operator:     f,
		WorkerChanel: make(chan struct{}, tc.WorkerNum),
		TimeOut:      tc.TimeOut,
	}
}

// Tasker ....
type Tasker interface {
	Add()
	taskOperator()
	Run()
}

// BuildTaskOperator ....
func (t *Task) taskOperator(f func(interface{}) error, v interface{}) {
	c := make(chan error)
	go func() {
		result := f(v)
		c <- result
	}()
	select {
	case <-c:
		log.Info("done")
	case <-time.After(t.TimeOut):
		log.Info("task timeout")
	}
	<-t.WorkerChanel
	t.Wg.Done()
}

// Add ....
func (t *Task) Add(i ...interface{}) {
	for _, v := range i {
		t.Args = append(t.Args, v)
	}
}

// Run ....
func (t *Task) Run() {
	log.Info("start")
	log.Info("total tasks: ", len(t.Args))
	log.Info("task build...")
	for _, v := range t.Args {
		t.Wg.Add(1)
		t.WorkerChanel <- struct{}{}
		go t.taskOperator(t.Operator, v)
	}
	t.Wg.Wait()
	log.Info("ok")
}
