package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// defer close(ch)

	go func() {
		// defer close(ch)
		for {
			ch <- 10
			time.Sleep(time.Second)
		}

	}()
	i := 0
	for v := range ch {
		fmt.Println(v)
		i++
	}
	fmt.Printf("%d\n", i)
}
