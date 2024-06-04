package adm

import "fmt"

var count int

func init() {
	count++
	fmt.Println("oo")
}

func Get() int {
	return count
}
