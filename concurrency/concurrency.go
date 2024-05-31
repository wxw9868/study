package concurrency

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// 查看协程数量
func CheckTheNumberOfGoroutines() {
	gonum := runtime.NumGoroutine()
	fmt.Printf("NumGoroutine: %d\n", gonum)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
			time.Sleep(time.Second)
		}()
	}
	gonum = runtime.NumGoroutine()
	fmt.Printf("NumGoroutine: %d\n", gonum)
	time.Sleep(time.Second * 2)
	gonum = runtime.NumGoroutine()
	fmt.Printf("NumGoroutine: %d\n", gonum)
}

// UseChannelsToControlOutputOrder 用 channel 控制输出顺序
func UseChannelsToControlOutputOrder() {
	var wg = sync.WaitGroup{}
	var ch = make(chan struct{}, 1)
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(ch chan struct{}, i int) {
			defer wg.Done()
			ch <- struct{}{}
			fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
			time.Sleep(time.Second)
			<-ch
		}(ch, i)
	}
	wg.Wait()
}

// TimeoutExit 超时退出
func TimeoutExit() {
	// WithTimeout控制
	withTimeout := func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		neverReady := make(chan struct{})
		go func() {
			time.Sleep(time.Second * 2)
			neverReady <- struct{}{}
		}()
		select {
		case <-neverReady:
			fmt.Println("ready")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		}
	}
	withTimeout()

	// 定时器控制
	timer := func() {
		timer := time.NewTimer(time.Second)
		for {
			select {
			case <-timer.C:
				fmt.Println("Timeout Exit")
				return
			default:
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Work running")
			}
		}
	}
	timer()
}

// ProduceAndConsume 生产者消费者
func ProduceAndConsume() {
	var wg sync.WaitGroup

	job := make(chan int)
	timeout := make(chan bool)

	go func() {
		for i := 0; ; i++ {
			select {
			case <-timeout:
				close(job)
				return
			default:
				job <- i
				fmt.Println("produce:", i)
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range job {
			fmt.Println("consume:", i)
		}
	}()

	time.Sleep(time.Second)
	timeout <- true

	wg.Wait()
}

// ManuallyExitTask 手动退出工作任务
func ManuallyExitTask() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for {
				select {
				case s := <-strings:
					// Do something interesting
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return completed
	}

	done := make(chan interface{})
	workString := make(chan string)
	terminated := doWork(done, workString)

	go func() {
		// 工作内容
		for i := 0; i < 5; i++ {
			workString <- "work " + strconv.Itoa(i)
		}

		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}

// ReusingGoroutine 重复使用协程
func ReusingGoroutine() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var ch = make(chan int)
	var data = make(map[int]int, 0)
	// data := new(sync.Map)
	// data := make([]int, 0)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
		E:
			for {
				r, ok := <-ch
				if ok {
					// 使用互斥锁保护 data
					mu.Lock()
					data[r] = r
					// data.LoadOrStore(r, r)
					// data = append(data, r)
					mu.Unlock()
				} else {
					fmt.Println(i, r, ok)
					break E
				}
			}
		}(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Wait()

	fmt.Println(data)
}

// UsingMutexLock 使用互斥锁
func UsingMutexLock() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var data = make([]int, 0)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// 使用互斥锁保护 data
			mu.Lock()
			data = append(data, i)
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(data)
}

// UsingMutexesToResolveDataRaces 使用互斥锁解决数据竞争
func UsingMutexesToResolveDataRaces() {
	var wg sync.WaitGroup
	var mux sync.Mutex
	slice := []int{1, 2, 3, 4, 5}
	ans := 0

	for _, v := range slice {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mux.Lock()
			ans += v
			mux.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("ans: %v\n", ans)
}

// AtomicOperations 原子级操作
func AtomicOperations() {
	var wg sync.WaitGroup
	ans := int64(0)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i *int64) {
			defer wg.Done()
			atomic.AddInt64(i, 1)
		}(&wg, &ans)
	}
	wg.Wait()
	fmt.Println(ans)
}
