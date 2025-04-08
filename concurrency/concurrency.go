package concurrency

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

// go如何用两个协程交替打印出123456
func PrintOddAndEven() {
	signalCh := make(chan struct{})
	oddCh := make(chan int)
	evenCh := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		timer := time.NewTimer(time.Second * 1)
		i := 0
		for {
			select {
			case <-timer.C:
				fmt.Printf("even: %s\n", "PUT END")
				signalCh <- struct{}{}
				return
			default:
				time.Sleep(time.Millisecond * 100)
				i++
				if i%2 != 0 {
					oddCh <- i
				} else {
					evenCh <- i
				}
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case odd, ok := <-oddCh: // 奇数
				if ok {
					fmt.Printf("odd: %d\n", odd)
				}
			case even, ok := <-evenCh: // 偶数
				if ok {
					fmt.Printf("even: %d\n", even)
				}
			case <-signalCh:
				fmt.Printf("even: %s\n", "OUT END")
				return
			}
		}
	}()
	wg.Wait()
}

// WorkDone 启动 2个goroutine 2秒后取消， 第一个协程1秒执行完，第二个协程3秒执行完。
func WorkDone() {
	doWork := func(done <-chan any, strings <-chan string) <-chan any {
		completed := make(chan any)
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

	done := make(chan any)
	terminated := doWork(done, nil)

	go func() {
		// Cancel the operation after 1 second.
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println("Done.")
}

// TimeoutCloseTask 超时关闭channel并退出程序
func TimeoutCloseTask() {
	jobs := make(chan int)
	timeout := make(chan bool)
	var wg sync.WaitGroup
	go func() {
		time.Sleep(time.Second * 3)
		timeout <- true
	}()
	go func() {
		for i := 0; ; i++ {
			select {
			case <-timeout:
				close(jobs)
				return
			default:
				jobs <- i
				fmt.Println("produce:", i)
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range jobs {
			fmt.Println("consume:", i)
		}
	}()
	wg.Wait()
}

// TimeoutOnCloseChannel 超时关闭channel并退出程序
func TimeoutOnCloseChannel() {
	wg := sync.WaitGroup{}
	c := make(chan struct{})
	for i := range 10 {
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

// MutexAndGoroutine 锁和goroutine的关系
func MutexAndGoroutine() {
	ch := make(chan struct{}, 2)
	var mutex sync.Mutex
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		fmt.Println("goroutine1: 我会锁定大概 2s")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1: 我解锁了，你们去抢吧")
		ch <- struct{}{}
	}()
	go func() {
		fmt.Println("goroutine2: 等待解锁")
		mutex.Lock()
		defer mutex.Unlock()
		fmt.Println("goroutine2: 欧耶，我也解锁了")
		ch <- struct{}{}
	}()
	// 等待 routine 执行结束
	for i := 0; i < 2; i++ {
		<-ch
	}
}

// 查看协程数量
func CheckTheNumberOfGoroutines() {
	gonum := runtime.NumGoroutine()
	fmt.Printf("NumGoroutine: %d\n", gonum)
	for i := range 5 {
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
	for i := range 8 {
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

// NewTicker 定时轮询
func NewTicker() {
	var wg sync.WaitGroup

	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan int)

	wg.Add(1)
	go func() {
		defer func() {
			fmt.Println("child routine bootstrap end")
			wg.Done()
		}()
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
	}()
	time.Sleep(10 * time.Second)

	quit <- 1
	wg.Wait()
}

// RunTask 手动停止定时任务
func RunTask() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	timeTask := func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		task := func(i int) {
			fmt.Println(i)
		}
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
	wg.Add(1)
	go timeTask(ctx, wg)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
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
		for i := range 5 {
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

	for i := range runtime.NumCPU() {
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
		for i := range 10 {
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

	for i := range 5 {
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

// UsingReadWriteLocks 使用读写锁
func UsingReadWriteLocks() {
	var count int
	var rw sync.RWMutex
	ch := make(chan struct{}, 10)
	read := func(n int, ch chan struct{}) {
		rw.RLock()
		fmt.Printf("routine %d 进入读操作...\n", n)
		v := count
		fmt.Printf("routine %d 读取结束，值为：%d\n", n, v)
		rw.RUnlock()
		ch <- struct{}{}
	}
	write := func(n int, ch chan struct{}) {
		rw.Lock()
		fmt.Printf("routine %d 进入写操作...\n", n)
		v := rand.Intn(1000)
		count = v
		fmt.Printf("routine %d 写入结束，新值为：%d\n", n, v)
		rw.Unlock()
		ch <- struct{}{}
	}
	for i := range 5 {
		go read(i, ch)
	}
	for i := range 5 {
		go write(i, ch)
	}
	for range 10 {
		<-ch
	}
}

// AtomicOperations 原子级操作
func AtomicOperations() {
	var wg sync.WaitGroup
	ans := int64(0)
	for range 10 {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i *int64) {
			defer wg.Done()
			atomic.AddInt64(i, 1)
		}(&wg, &ans)
	}
	wg.Wait()
	fmt.Println(ans)
}
