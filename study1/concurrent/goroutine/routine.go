package routine

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// GoDone 启动 2个goroutine 2秒后取消， 第一个协程1秒执行完，第二个协程3秒执行完。
func GoDone() {
	wg.Add(2)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	go work(ctx, time.Second, "go1")
	go work(ctx, time.Second*3, "go2")

	wg.Wait()
	cancel()
}

func work(ctx context.Context, t time.Duration, msg string) {
	defer wg.Done()
	timer := time.NewTimer(t)
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			fmt.Println(msg)
			return
		}
	}
}

func DoOrder() {
	ch1, ch2, ch3, ch4 := make(chan struct{}), make(chan struct{}), make(chan struct{}), make(chan struct{})
	wg := new(sync.WaitGroup)
	wg.Add(5)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("1")
		ch1 <- struct{}{}
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		msg := <-ch1
		fmt.Println("2")
		ch2 <- msg
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		msg := <-ch2
		fmt.Println("3")
		ch3 <- msg
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		msg := <-ch3
		fmt.Println("4")
		ch4 <- msg
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		<-ch4
		fmt.Println("5")
	}(wg)

	wg.Wait()
}

func DoOrder1() {
	ch1, ch2, ch3, ch4, ch5 := make(chan struct{}), make(chan struct{}), make(chan struct{}), make(chan struct{}), make(chan struct{})

	go func() {
		fmt.Println("1")
		ch1 <- struct{}{}
	}()

	go func() {
		msg := <-ch1
		fmt.Println("2")
		ch2 <- msg
	}()

	go func() {
		msg := <-ch2
		fmt.Println("3")
		ch3 <- msg
	}()

	go func() {
		msg := <-ch3
		fmt.Println("4")
		ch4 <- msg
	}()

	go func() {
		msg := <-ch4
		fmt.Println("5")
		time.Sleep(time.Second)
		ch5 <- msg
	}()

	timer := time.NewTimer(time.Millisecond * 100)
	for {
		select {
		case <-ch5:
			close(ch1)
			close(ch2)
			close(ch3)
			close(ch4)
			close(ch5)
			fmt.Println("end")
			return
		case <-timer.C:
			fmt.Println("time out")
			return
		}
	}
}

func PutMode() {
	ch := make(chan int, 10)
	putMode := 0
	for i := 0; i < 1000; i++ {
		ch <- i
		wg.Add(1)
		go AFunc(ch, putMode)
	}
	wg.Wait()
	close(ch)
}

func AFunc(ch chan int, putMode int) {
	val := <-ch
	switch putMode {
	case 0:
		fmt.Printf("Vaule=%d\n", val)
	case 1:
		fmt.Printf("Vaule=%d\n", val)
		for i := 1; i <= 5; i++ {
			ch <- i * val
		}
	case 2:
		fmt.Printf("Vaule=%d\n", val)
		for i := 1; i <= 5; i++ {
			<-ch
		}
	}
	wg.Done()
	fmt.Println("WaitGroup Done", val)
}

// WaitGroup2 使用wg等待所有routine执行完毕，并输出相应的提示信息
func WaitGroup2() {
	ch := make(chan int)
	execMode := 0
	switch execMode {
	case 0:
		go Afunc(ch)
		ch <- 100
	case 1:
		ch <- 100
		go Afunc(ch)
	}
	wg.Wait()
	close(ch)
}

func Afunc(ch chan int) {
	wg.Add(1)
FlAG:
	for {
		select {
		case val := <-ch:
			fmt.Println(val)
			break FlAG
		}
	}
	wg.Done()
	fmt.Println("WaitGroup Done")
}

// WaitGroup 使用wg等待所有routine执行完毕，并输出相应的提示信息
func WaitGroup() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Afunction(i)
	}
	wg.Wait()
}

func Afunction(index int) {
	fmt.Println(index)
	wg.Done()
}

// 计数器
var count int

// 创建互斥对象，通过互斥对象加锁保护count变量，同一时间只能一个协程访问
var mu sync.Mutex

func Count() {
	for i := 0; i < 2; i++ {
		// 循环创建两个协程，执行匿名函数
		go func() {
			// 加锁
			mu.Lock()
			// 延迟释放锁 - 匿名函数执行结束才会释放锁
			defer mu.Unlock()

			// 下面的代码同一时间只有一个协程在运行
			// 对count累加计数
			for j := 0; j < 100; j++ {
				count++
			}
		}()
	}

	// 先休眠5秒，等前面的协程执行结束
	time.Sleep(5 * time.Second)

	// 打印计数值
	fmt.Println(count)
}
