package main

import (
	"fmt"
	"study/basics/init/a"
	"study/basics/init/b"
	"study/basics/init/c"
)

// 查看 包被多次调用时，包中 init() 是否多次执行？
func main() {
	a.Print()
	b.Print()
	fmt.Println(c.Get())
}
