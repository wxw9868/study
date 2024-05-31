package basics

import (
	"encoding/json"
	"fmt"
	"log"
)

// 基础
func Basics() {
	fmt.Println("数值原样输出: ")
	i := -5
	j := +5
	fmt.Printf("%+d %+d\n", i, j)

	fmt.Println("数组和切片: ")
	arr := [3]int{1, 2, 3}
	fmt.Printf("arr: p=%p,len=%d,cap=%d\n", &arr, len(arr), cap(arr))
	fmt.Printf("arr: p=%p\n", &arr[0])
	fmt.Printf("arr: p=%p\n", &arr[1])
	fmt.Printf("arr: p=%p\n", &arr[2])
	s1 := []int{1, 2, 3}
	fmt.Printf("s1: p=%p\n", &s1[0])
	fmt.Printf("s1: p=%p\n", &s1[1])
	fmt.Printf("s1: p=%p\n", &s1[2])
	fmt.Printf("s1: p=%p,len=%d,cap=%d\n", &s1, len(s1), cap(s1))
	s2 := []int{4, 5}
	fmt.Printf("s2: p=%p,len=%d,cap=%d\n", &s2, len(s2), cap(s2))
	s1 = append(s1, s2...)
	fmt.Println(s1)
	fmt.Printf("s1: p=%p,len=%d,cap=%d\n", &s1, len(s1), cap(s1))
}

// 函数和通道不能转json
func StructToJson() {
	type S struct {
		A int           `json:"a"`
		B *int          `json:"b"`
		C float64       `json:"c"`
		D func() string `json:"d"`
		E chan struct{} `json:"e"`
	}

	s := S{
		A: 1,
		B: nil,
		C: 12.15,
		D: func() string {
			return "NowCoder"
		},
		E: make(chan struct{}),
	}

	_, err := json.Marshal(s)
	if err != nil {
		log.Printf("err occurred... %s\n", err)
		return
	}
	log.Printf("everything is ok.")
}
