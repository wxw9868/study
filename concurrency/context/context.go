package ctx

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Demo1() {
	//gen函数在新的goroutine内产生整数，并发送这些整数到返回的无缓冲通道内
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done(): //阻塞直到ctx的Done通道被取消
					fmt.Println(ctx.Err()) //ctx被取消后，打印错误信息
					return
				case dst <- n: //向返回的通道写入数据
					n++ //递增数据
				}
			}
		}()
		return dst
	}

	//创建一个可以被取消的context
	ctx, cancel := context.WithCancel(context.Background())
	//从gen函数返回的通道内读取5个元素并打印
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel() //如果当了第5个元素则取消ctx，则ctx的Done通道会被取消并返回0值
			break
		}
	}
	//main函数所在goroutine休眠1s
	time.Sleep(1 * time.Second)
	cancel()
}

func Demo2() {
	//1.截止时间为当前时间加上100ms
	d := time.Now().Add(100 * time.Millisecond)
	//2.基于空context创建一个deadline上下文
	ctx, cancel := context.WithDeadline(context.Background(), d)
	//3.等main函数结束后，上下文会被关闭，如果不及时关闭会导致上下文ctx和其父context存在的周期比我们想要的长
	defer cancel()
	//4.select块
	select {
	case <-time.After(1 * time.Second): //4.1如果1s上下文还没被取消，则超时打印
		fmt.Println("overslept")
	case <-ctx.Done(): //4.2如果上下文ctx在1s内被取消或者超时了，则打印错误
		fmt.Println(ctx.Err())
	}
}

func Demo3() {
	var wg sync.WaitGroup
	wg.Add(2)
	//自定义类型
	type String string
	//创建函数，内部从ctx内获取变量k的值并打印
	f := func(ctx context.Context, k String) {
		defer wg.Done()
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	//创建k,并关联上下文
	k := String("language")
	ctx := context.WithValue(context.Background(), k, "Go")
	//开启goroutine
	go f(ctx, k)
	go f(ctx, "color")
	wg.Wait()
}

func Demo4() {
	fmt.Println("run demo")
	ctx, cancel := context.WithTimeout(context.Background(), 9*time.Second)
	go watch(ctx, "[线程1]")
	go watch(ctx, "[线程2]")
	go watch(ctx, "[线程3]")

	index := 0
	for {
		index++
		fmt.Printf("%d 秒过去了 \n", index)
		time.Sleep(1 * time.Second)
		if index > 10 {
			break
		}
	}

	fmt.Println("通知停止监控")
	// 其实此时已经超时，协程已经提前退出
	cancel()

	// 防止主进程提前退出
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s 监控退出，停止了...\n", name)
			return
		default:
			fmt.Printf("%s goroutine监控中... \n", name)
			time.Sleep(2 * time.Second)
		}
	}
}
