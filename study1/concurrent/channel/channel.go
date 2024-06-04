package ch

import (
	"fmt"
	"time"
)

// Study1 在通道中写入一个值后关闭，读取两次，分别会得到的值是什么？
func Study1() {
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	go func() {
		s, ok := <-ch
		fmt.Println("s1: ", s, ok)
	}()

	go func() {
		s, ok := <-ch
		fmt.Println("s2: ", s, ok)
	}()

	time.Sleep(time.Millisecond)
}

// Study2 如何判断channel是否关闭
func Study2() {
	ch := make(chan int)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("panic: ", r)
			}
		}()
		ch <- 1
		//close(channel)
		ch <- 2
	}()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("panic: ", r)
			}
		}()
		//此时如果 channel 关闭，ok 值为 false，如果 channel 没有关闭，则会漏掉一个 jobs
		r, ok := <-ch
		if !ok {
			fmt.Println("close")
		}
		ch <- r
	}()

	go func() {
		s, ok := <-ch
		fmt.Println("s1: ", s, ok)
	}()

	go func() {
		s, ok := <-ch
		fmt.Println("s2: ", s, ok)
	}()

	time.Sleep(time.Millisecond)
}

// Study3 关闭通道后会输出什么结果？
func Study3() {
	ch := make(chan int)

	go func() {
		close(ch)
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 1000)
			s, ok := <-ch
			fmt.Println("s3: ", s, ok)
		}
	}()

	time.Sleep(time.Millisecond * 10000)
}

// Study4 关闭通道后继续给通道写入数据会出现什么结果？
func Study4() {
	ch := make(chan int)

	close(ch)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("panic: ", r)
			}
		}()

		ch <- 1
	}()

	time.Sleep(time.Millisecond * 100)
}

func Study5() {
	//channel := make(chan int)
	//channel <- 1

	ch1 := make(chan int, 10)
	ch1 <- 1
}

func GenerateNumber() {
	generator := xrange()

	for i := 0; i < 1000; i++ {
		fmt.Println(<-generator)
	}
}

func xrange() chan int {
	var ch chan int = make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}
