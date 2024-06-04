package other

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	m := NewMap()

	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU())

	f := func() {
		defer wg.Done()
		m.Set("key", "go")
		m.Set("name", []int{1, 2})
	}
	for i := 0; i < runtime.NumCPU(); i++ {
		go f()
	}

	wg.Wait()
	m.Get("key")
	fmt.Println(m.Len())
	fmt.Println(m.Print())
}

var m sync.Map

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
}

func BenchmarkName1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.Load(i)
	}
}

func BenchmarkD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		do()
	}
}

type Config struct{}

func (c *Config) init(filename string) {
	fmt.Printf("mock [%s] config initial done!\n", filename)
}

var (
	//once sync.Once
	once Once
	cfg  *Config
)

func do() {
	cfg = &Config{}

	go once.Do(func() {
		time.Sleep(3 * time.Second)
		cfg.init("first file path")
	})

	time.Sleep(time.Second)
	once.Do(func() {
		fmt.Println("22")
		time.Sleep(time.Second)
		cfg.init("second file path")
	})
	fmt.Println("运行到这里！")
	time.Sleep(5 * time.Second)
}
