package main

import (
	"fmt"
	"sync"
)

// 交替打印数字和字母
// 使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一个 goroutine 打印字母， 最终效果如下：
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
func AlternatePrinting() {
	// for i := 65; i < 91; i++ {
	// 	fmt.Print(string(byte(i)))
	// }
	// fmt.Println()
	// fmt.Println(string(byte(65)), string(byte(90)))  // A - Z
	// fmt.Println(string(byte(97)), string(byte(122))) // a - z

	var zimu = 64
	var shuzi = 0
	var wg sync.WaitGroup
	var s1 = make(chan struct{})
	var s2 = make(chan struct{})

	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			<-s1
			shuzi++
			fmt.Print(shuzi)
			shuzi++
			fmt.Print(shuzi)
			if shuzi == 28 {
				fmt.Println()
				return
			}
			s2 <- struct{}{}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			if string(byte(zimu)) == "Z" {
				return
			}
			<-s2
			zimu++
			fmt.Print(string(byte(zimu)))
			zimu++
			fmt.Print(string(byte(zimu)))
			s1 <- struct{}{}
		}
	}()

	s1 <- struct{}{}
	wg.Wait()
	close(s1)
	close(s2)
}

// 问题描述
// 请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
// 给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。
func ReverseString(s string) string {
	if len(s) > 5000 {
		return "字符串的长度小于等于5000"
	}

	var str string
	for i := len(s) - 1; i >= 0; i-- {
		str += string(s[i])
	}
	return str
}

func f(n int) (r int) {
	defer func() {
		r += n
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var f func()

	f()
	f = func() {
		r += 2
	}

	return n + 1
}

func main() {
	var wg sync.WaitGroup
	wg.Done()
	wg.Wait()
}

// func main() {
// 	gonum := runtime.NumGoroutine()
// 	fmt.Printf("NumGoroutine: %d\n", gonum)
// 	go func() {
// 		time.Sleep(time.Second)
// 		fmt.Println("1")
// 	}()

// 	go func() {
// 		time.Sleep(time.Second)
// 		fmt.Println("2")
// 	}()

// 	gonum = runtime.NumGoroutine()
// 	fmt.Printf("NumGoroutine: %d\n", gonum)
// 	time.Sleep(time.Second * 2)
// 	gonum = runtime.NumGoroutine()
// 	fmt.Printf("NumGoroutine: %d\n", gonum)

// 	i := -5
// 	j := +5
// 	fmt.Printf("%+d %+d", i, j)

// 	arr := [3]int{1, 2, 3}
// 	fmt.Printf("arr: p=%p,len=%d,cap=%d\n", &arr, len(arr), cap(arr))
// 	fmt.Printf("arr: p=%p\n", &arr[0])
// 	fmt.Printf("arr: p=%p\n", &arr[1])
// 	fmt.Printf("arr: p=%p\n", &arr[2])
// 	s1 := []int{1, 2, 3}
// 	s2 := []int{4, 5}
// 	fmt.Printf("p=%p\n", &s1[0])
// 	fmt.Printf("p=%p\n", &s1[1])
// 	fmt.Printf("p=%p\n", &s1[2])
// 	fmt.Printf("s1: p=%p,len=%d,cap=%d\n", &s1, len(s1), cap(s1))
// 	fmt.Printf("s2: p=%p,len=%d,cap=%d\n", &s2, len(s2), cap(s2))
// 	s1 = append(s1, s2...)
// 	fmt.Printf("s1: p=%p,len=%d,cap=%d\n", &s1, len(s1), cap(s1))
// 	fmt.Printf("s2: p=%p,len=%d,cap=%d\n", &s2, len(s2), cap(s2))
// 	fmt.Println(s1)

// 	s := S{
// 		A: 1,
// 		B: nil,
// 		C: 12.15,
// 		D: func() string {
// 			return "NowCoder"
// 		},
// 		E: make(chan struct{}),
// 	}
// 	_, err := json.Marshal(&s)
// 	if err != nil {
// 		log.Printf("err occurred.. %v", err)
// 		return
// 	}
// 	log.Printf("everything is ok.")
// }

// type S struct {
// 	A int
// 	B *int
// 	C float64
// 	D func() string
// 	E chan struct{}
// }
