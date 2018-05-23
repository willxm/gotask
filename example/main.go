package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/willxm/gotask"
)

var test gotask.TaskHandle = func(v gotask.TaskArg) error {
	// fmt.Println(v)
	time.Sleep(time.Duration(rand.Float64()*1000) * time.Millisecond)
	// time.Sleep(time.Second)
	fmt.Println(v)
	return nil
}

func main() {
	tf := gotask.TaskConfig{
		WorkerNum: 8,
		// defalt timeout is 30s
		// Timeout:   1000 * time.Millisecond,
	}

	task := tf.NewTask(test)

	for i := 0; i < 100; i++ {
		task.Add(i)
	}

	task.Run()
}
