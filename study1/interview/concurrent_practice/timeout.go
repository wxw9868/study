package concurrent_practice

import (
	"fmt"
	"sync"
	"time"
)

// TimeoutOnCloseChannel 超时关闭channel并退出程序
func TimeoutOnCloseChannel() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, close <-chan struct{}) {
			defer wg.Done()
			<-close
			fmt.Println(num)
		}(i, c)
	}

	if WaitTimeout(&wg, time.Second*5) {
		close(c)
		fmt.Println("timeout exit")
	}
	time.Sleep(time.Second * 10)
}

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	t := time.NewTimer(timeout)
	ch := make(chan bool)
	go func() {
		<-t.C
		ch <- true
	}()

	go func() {
		wg.Wait()
		ch <- false
	}()

	return <-ch
}
