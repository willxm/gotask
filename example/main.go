package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/willxm/gotask"
)

func test(v interface{}) error {
	// fmt.Println(v)
	time.Sleep(time.Duration(rand.Float64()*1000) * time.Millisecond)
	// time.Sleep(time.Second)
	fmt.Println(v)
	return nil
}

func main() {
	tf := gotask.TaskConfig{
		WorkerNum: 2,
		TimeOut:   500 * time.Millisecond,
	}

	task := tf.NewTask(test)

	for i := 0; i < 100; i++ {
		task.Add(i)
	}

	task.Run()

}
