package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup

	ans := int64(0)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go newGoRoutine(&wg, &ans)
	}
	wg.Wait()
	fmt.Println(ans)
}

func newGoRoutine(wg *sync.WaitGroup, i *int64) {
	defer wg.Done()
	atomic.AddInt64(i, 1)
	return
}
