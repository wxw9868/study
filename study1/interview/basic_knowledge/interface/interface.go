package iface

import "fmt"

func test(i interface{}) {
	r, ok := i.(func() (string, error))
	fmt.Println(ok)
	fmt.Println(r())
}

func New() (string, error) {
	return "wxw", nil
}

func Run() {
	test(New)
}
