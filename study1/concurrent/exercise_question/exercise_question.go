package exercise_question

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var (
	lock                 sync.Mutex
	myMap                = make(map[int]int)
	wg                   sync.WaitGroup
	catCH, fishCH, dogCH = make(chan struct{}, 100), make(chan struct{}), make(chan struct{})
)

// DanGoroutine 单协程写切片
func DanGoroutine() {
	s := time.Now()
	slice := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		//slice = append(slice, i)
		slice[i] = i
	}
	fmt.Println(len(slice))
	d := time.Since(s)
	fmt.Println(d)
}

// DuoGoroutine 多协程写切片
func DuoGoroutine() {
	s := time.Now()

	num := runtime.NumCPU()
	slice := make([]int, 10000000)
	ch := make(chan int, num)
	var wg sync.WaitGroup
	var mu sync.Mutex

	go func() {
		for i := 0; i < 10000000; i++ {
			ch <- i
		}
		close(ch)
	}()

	f := func() {
		defer wg.Done()
		defer mu.Unlock()
		mu.Lock()
		for i := range ch {
			//slice = append(slice, i)
			slice[i] = i
		}
	}

	wg.Add(num)
	for i := 0; i < num; i++ {
		go f()
	}

	wg.Wait()
	fmt.Println(len(slice))

	d := time.Since(s)
	fmt.Println(d)
	fmt.Println(runtime.NumCPU())
}

func G() {
	s := time.Now()

	num := runtime.NumCPU()
	slice := make([]int, 10000)
	ch := make(chan int, num)
	var wg sync.WaitGroup
	//var mu sync.Mutex

	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(i int) {
			ch <- i
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		slice[i] = i
	}

	fmt.Println(len(slice))
	d := time.Since(s)
	fmt.Println(d)
}

// DumpStr
// 交叉打印数字和字母
// 结果：12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func DumpStr() {
	var strCh = make(chan string, 2)
	var numCh = make(chan string, 2)
	var wg sync.WaitGroup

	s := ""
	var str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	wg.Add(2)
	go func() {
		for i := 1; i < 28; i = i + 2 {
			if i > 26 {
				wg.Done()
			}
			numCh <- strconv.Itoa(i)
			numCh <- strconv.Itoa(i + 1)
			s += <-strCh
			s += <-strCh

		}
	}()
	go func() {
		for i := 0; i < 28; i = i + 2 {
			s += <-numCh
			s += <-numCh
			if i > 25 {
				strCh <- ""
				strCh <- ""
				wg.Done()
			} else {
				strCh <- string(str[i])
				strCh <- string(str[i+1])
			}
		}
	}()

	wg.Wait()
	fmt.Println(s)
}

// Order （Goroutine）有三个函数，分别打印"cat", "fish","dog"要求每一个函数都用一个goroutine，按照顺序打印100次。
// 此题目考察channel，用三个无缓冲channel，如果一个channel收到信号则通知下一个。
func Order() {
	wg.Add(300)
	for i := 0; i < 100; i++ {
		go cat()
		go fish()
		go dog()
		catCH <- struct{}{}
	}
	wg.Wait()
}

func cat() {
	defer wg.Done()
	<-catCH
	fmt.Println("cat")
	fishCH <- struct{}{}
}

func fish() {
	defer wg.Done()
	<-fishCH
	fmt.Println("fish")
	dogCH <- struct{}{}
}

func dog() {
	defer wg.Done()
	<-dogCH
	fmt.Println("dog")
}

// Factorial
// 需求:现在要计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中。
// 最后显示出来。要求使用goroutine完成
func Factorial() {
	wg.Add(20)
	for i := 1; i <= 20; i++ {
		go doFactorial(i)
	}

	wg.Wait()
	//time.Sleep(time.Second * 10)
	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("k: %d, v: %d\n", k, v)
	}
	lock.Unlock()
}

func doFactorial(n int) {
	defer wg.Done()
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock()
	myMap[n] = res
	lock.Unlock()
}

// AddSlice 错误的编写方式
func AddSlice() {
	var a []int

	for i := 0; i < 10000; i++ {
		go func() {
			a = append(a, 1) // 多协程并发读写slice
		}()
	}

	fmt.Println(len(a))
}

// AddSliceMutex 第一种方式
func AddSliceMutex() {
	num := 10000

	var a []int
	var l sync.Mutex

	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		b := i
		go func() {
			l.Lock()
			a = append(a, b)
			l.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(len(a))
}

// AddSliceWait 第三种方式
func AddSliceWait() {
	num := 10000

	var a []int
	//a := make([]int, num, num)

	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		i := i //必须使用局部变量
		go func() {
			a[i] = 1
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(len(a))
	//count := 0
	//for i := range a {
	//	if a[i] != 0 {
	//		count++
	//	}
	//}
	//fmt.Println(count)
}

func DoType() {
	timer := time.NewTimer(time.Second * 5)
	slice := []int{23, 32, 78, 43, 76, 65, 345, 762, 676, 915, 86, 68, 96, 10, 22}
	sliceLen := len(slice)
	size := 3
	tag := 345

	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan struct{})
	for i := 0; i < sliceLen; i += size {
		end := i + size
		if end >= sliceLen {
			end = sliceLen - 1
		}
		go do(ctx, slice[i:end], tag, resultChan)
	}

	select {
	case <-timer.C:
		_, _ = fmt.Fprintln(os.Stderr, "Timeout! Not Found")
		cancel()
	case <-resultChan:
		_, _ = fmt.Fprintf(os.Stdout, "Found it!\n")
		cancel()
	}

	time.Sleep(time.Second * 2)
}

func do(ctx context.Context, data []int, tag int, resultChan chan struct{}) {
	for _, v := range data {
		select {
		case <-ctx.Done():
			_, _ = fmt.Fprintf(os.Stdout, "Task cancelded! \n")
			return
		default:
		}

		_, _ = fmt.Fprintf(os.Stdout, "v: %d \n", v)
		time.Sleep(time.Millisecond * 1500)
		if v == tag {
			resultChan <- struct{}{}
			return
		}
	}
}
