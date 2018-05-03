package main

import (
	"fmt"
	"time"

	"github.com/willxm/gotask"
)

func test(v interface{}) {
	fmt.Println(v)
	time.Sleep(time.Second)
}

func main() {
	tf := gotask.TaskConfig{
		WorkerNum: 6,
	}

	task := tf.NewTask(test)

	for i := 0; i < 100; i++ {
		task.Add(i)
	}

	task.Run()

}
