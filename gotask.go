package gotask

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	log "github.com/lytics/logrus"
)

/*
task func must be is type
func(interface{}) error
*/

type TaskArg interface{}
type TaskHandle func(TaskArg) error

// TaskConfig ....
type TaskConfig struct {
	Handle    func(interface{})
	WorkerNum int
	Timeout   time.Duration
}

// Task ....
type Task struct {
	Operator     func(TaskArg) error
	Args         []interface{}
	WorkerChanel chan struct{}
	Wg           *sync.WaitGroup
	Timeout      time.Duration
}

// NewTask ....
func (tc *TaskConfig) NewTask(f TaskHandle) *Task {
	// defalt timeout 30s
	if tc.Timeout == 0 {
		tc.Timeout = 30 * time.Second
	}
	if tc.WorkerNum == 0 {
		tc.WorkerNum = runtime.NumCPU()
	}
	return &Task{
		Wg:           &sync.WaitGroup{},
		Operator:     f,
		WorkerChanel: make(chan struct{}, tc.WorkerNum),
		Timeout:      tc.Timeout,
	}
}

// Tasker ....
type Tasker interface {
	Add()
	taskOperator()
	Run()
}

// BuildTaskOperator ....
func (t *Task) taskOperator(f TaskHandle, v TaskArg) {
	c := make(chan error)
	go func() {
		result := f(v)
		c <- result
	}()
	select {
	case <-c:
		log.Info("done")
	case <-time.After(t.Timeout):
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
	for k, v := range t.Args {
		t.Wg.Add(1)
		t.WorkerChanel <- struct{}{}
		go t.taskOperator(t.Operator, v)
		fmt.Printf("\f%s%d%%", Bar(k, 20), k)
	}
	t.Wg.Wait()
	log.Info("ok")
}
