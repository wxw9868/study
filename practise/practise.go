package practise

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Test struct {
	name string
}

func (t *Test) Point() { // t 为指针
	fmt.Println(t.name)
}

func Point() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.Point() //输出 c c c
	}
}

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

func StudyContext() {
	var wg sync.WaitGroup
	ctx := context.WithValue(context.Background(), "key", "value")
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	wg.Add(1)
	go coroutine(ctx, &wg)
	wg.Wait()
	cancel()
}

func coroutine(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
E:
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s\n", "coroutine exit")
			break E
		default:
			value := ctx.Value("key")
			fmt.Println(value)
			time.Sleep(time.Millisecond * 100)
		}
	}
}
