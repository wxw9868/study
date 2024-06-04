package concurrent_practice

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// NewTicker 定时轮询
func NewTicker() {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("child routine bootstrap start")
		for {
			select {
			case <-ticker.C:
				fmt.Println("ticker .")
			case <-quit:
				fmt.Println("work well .")
				ticker.Stop()
				return
			}
		}
		fmt.Println("child routine bootstrap end")
	}()
	time.Sleep(10 * time.Second)
	quit <- 1
	wg.Wait()
}

func RunTask() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go timeTask(ctx, wg)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}

func timeTask(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)
	var i int
	for {
		select {
		case <-ticker.C:
			i++
			task(i)
		case <-ctx.Done():
			fmt.Println("Done")
			return
		}
	}
}

func task(i int) {
	fmt.Println(i)
}
