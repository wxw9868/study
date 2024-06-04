package singleton

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	ins1 := GetInstance()
	ins2 := GetInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}

func TestDofunc(t *testing.T) {
	o := new(Once)
	for i := 0; i < 1000; i++ {
		go func() {
			o.Do(func() {
				fmt.Println("1")
			})
		}()
	}

	time.Sleep(3 * time.Second)
}
