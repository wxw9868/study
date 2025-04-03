package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	intSlice := []int{1, 2, 3, 4, 5}
	wg.Add(len(intSlice))
	ans1, ans2 := 0, 0
	for _, v := range intSlice {
		vv := v
		go func() {
			defer wg.Done()
			ans1 += v
			ans2 += vv
		}()
	}
	wg.Wait()
	fmt.Printf("ans1:%v,ans2:%v", ans1, ans2)
}
