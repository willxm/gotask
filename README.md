# gotask
a mini batch task processing toolkit

---
## install

```shell
$ go get github.com/willxm/gotask
```

## usage

```golang
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
		WorkerNum: 10,
	}

	task := tf.NewTask(test)

	for i := 0; i < 100; i++ {
		task.Add(i)
	}

	task.Run()

}
```

more example in /example/main.go
