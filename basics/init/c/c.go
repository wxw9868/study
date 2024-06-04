package c

import "fmt"

var count int

func init() {
	count++
	fmt.Println("init")
}

func Get() int {
	return count
}
