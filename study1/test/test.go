package test

import (
	"fmt"
)

var k = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// XH 循环
func XH() {
	for i := 0; i < len(k); i++ {
		for j := i + 1; j < len(k); j++ {
			if k[i] < k[j] {
				k[i], k[j] = k[j], k[i]
			}
		}
	}
	fmt.Println(k[0], k[1])
}

// DG 递归
func DG(a int) {
	i := a
	for j := i + 1; j < len(k); j++ {
		if k[i] < k[j] {
			k[i], k[j] = k[j], k[i]
		}
	}
	if i >= len(k)-1 {
		fmt.Println(k[0], k[1])
		return
	}
	i++
	DG(i)
}
