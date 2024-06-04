package pool

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	//创建一个Task
	f := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	//创建一个协程池,最大开启3个协程worker
	c := runtime.NumCPU()
	p := NewPool(c)

	//开一个协程 不断的向 Pool 输送打印一条时间的task任务
	go func() {
		for {
			p.EntryChannel <- f
		}
	}()

	//启动协程池p
	p.Run()
}
