package main

import (
	"fmt"
	"github.com/wxw9868/study/basics/init/a"
	"github.com/wxw9868/study/basics/init/b"
	"github.com/wxw9868/study/basics/init/c"
)

// 查看 包被多次调用时，包中 init() 是否多次执行？
func main() {
	a.Print()
	b.Print()
	fmt.Println(c.Get())
}
