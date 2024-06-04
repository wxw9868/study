package def

import (
	"fmt"
	"time"
)

func Demo() int {
	i := 0
	defer func() {
		fmt.Println("defer1")
	}()
	defer func() {
		i += 1
		fmt.Println("defer2")
	}()
	return i
}

func Return() int {
	i := 0
	defer func(i int) {
		fmt.Println("aaa")
		i++
		fmt.Println(i)
	}(i)
	return i + 2
}

func Defer() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("异常捕获: %s\n", r)
		}
	}()
	panic("panic")
}

func Run() {
	i, p := a()
	fmt.Println("return:", i, p, time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("c")

	defer func() { // 必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("d")
		if err := recover(); err != nil {
			fmt.Println(err) // 这里的err其实就是panic传入的内容
		}
		fmt.Println("e")
	}()

	f() //开始调用f

	fmt.Println("f") //这里开始下面代码不会再执行
}

func f() {
	fmt.Println("a")
	panic("异常信息")
	fmt.Println("b") //这里开始下面代码不会再执行
}

func a() (i int, p *int) {
	defer func(i int) {
		fmt.Println("defer3:", i, &i, time.Now().Format("2006-01-02 15:04:05"))
	}(i)
	defer fmt.Println("defer2:", i, &i, time.Now().Format("2006-01-02 15:04:05"))
	defer func() {
		fmt.Println("defer1:", i, &i, time.Now().Format("2006-01-02 15:04:05"))
	}()
	i++
	func() {
		fmt.Println("print1:", i, &i, time.Now().Format("2006-01-02 15:04:05"))
	}()
	fmt.Println("print2:", i, &i, time.Now().Format("2006-01-02 15:04:05"))
	return i, &i
}
