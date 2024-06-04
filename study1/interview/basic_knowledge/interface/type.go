package iface

import "fmt"

//需要传递函数
func callback(i int) {
	fmt.Println("i am callBack")
	fmt.Println(i)
}

// One main中调用的函数
func One(i int, f func(int)) {
	two(i, fun(f))
}

// one()中调用的函数
func two(i int, c Call) {
	c.call(i)
}

// 定义的type函数
type fun func(int)

// fun实现的Call接口的call()函数
func (f fun) call(i int) {
	f(i)
}

// Call 接口
type Call interface {
	call(int)
}
