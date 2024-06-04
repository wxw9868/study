package pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestName(t *testing.T) {
	nw := NewWorker(10, 20)
	dt := new(DoTask)
	for i := 0; i < 10; i++ {
		nw.Push(dt)
	}
	if err := nw.Close(); err != nil {
		fmt.Println("ERR: ", err)
	}
}

func BenchmarkDoTask_Do(b *testing.B) {
	//runtime.GC()
	//b.ReportAllocs()
	//b.ResetTimer()

	nw := NewWorker(10, 20)
	dt := new(DoTask)
	for i := 0; i < b.N; i++ {
		nw.Push(dt)
	}
	if err := nw.Close(); err != nil {
		fmt.Println("ERR: ", err)
	}
	//输出内存信息
	//var memStats runtime.MemStats
	//runtime.ReadMemStats(&memStats)
	//b.ReportMetric(float64(memStats.HeapInuse)/(1024*1024), "heap(MB)")
	//b.ReportMetric(float64(memStats.StackInuse)/(1024*1024), "stack(MB)")
	//b.ReportMetric(float64(memStats.StackInuse+memStats.HeapInuse)/(1024*1024), "total(MB)")
}

type testTask struct {
	wg *sync.WaitGroup
	ch chan struct{}
	m  bool
	p  int
}

func (t *testTask) Do() {
	if t.m {
		<-t.ch
	}
	t.wg.Done()
}

func newTestTask(wg *sync.WaitGroup) *testTask {
	return &testTask{
		wg: wg,
	}
}

func newMemTask(ch chan struct{}, wg *sync.WaitGroup) *testTask {
	return &testTask{
		wg: wg,
		ch: ch,
		m:  true,
	}
}

func (t *testTask) Priority() int {
	return t.p
}

func newPriorityTask(p int, wg *sync.WaitGroup) *testTask {
	return &testTask{
		wg: wg,
		p:  p,
	}
}

func runTest(b *testing.B, f func(i int, wg *sync.WaitGroup)) *sync.WaitGroup {
	//初始化
	runtime.GC()
	b.ReportAllocs()
	b.ResetTimer()
	var memStats1 runtime.MemStats
	runtime.ReadMemStats(&memStats1)
	wg := &sync.WaitGroup{}
	wg.Add(b.N)

	//执行测试
	for i := 0; i < b.N; i++ {
		f(i, wg)
	}

	//输出内存信息
	var memStats2 runtime.MemStats
	runtime.ReadMemStats(&memStats2)
	b.ReportMetric(float64(memStats2.HeapInuse-memStats1.HeapInuse)/(1024*1024), "heap(MB)")
	b.ReportMetric(float64(memStats2.StackInuse-memStats1.StackInuse)/(1024*1024), "stack(MB)")
	return wg
}

func BenchmarkNoGoroutine(b *testing.B) {
	wg := runTest(b, func(i int, wg *sync.WaitGroup) {
		t := newTestTask(wg)
		t.Do()
	})
	wg.Wait()
	b.StopTimer()
}

const (
	priority = 2
	number   = 2 * priority
)

func BenchmarkWorker(b *testing.B) {
	w := NewWorker(number, b.N)
	wg := runTest(b, func(i int, wg *sync.WaitGroup) {
		w.Push(newTestTask(wg))
	})
	w.Close()
	wg.Wait()
	b.StopTimer()
}

func BenchmarkPriorityWorker(b *testing.B) {
	w := NewPriorityWorker(number, b.N, priority)
	wg := runTest(b, func(i int, wg *sync.WaitGroup) {
		w.Push(newPriorityTask(i%priority, wg))
	})
	w.Close()
	wg.Wait()
	b.StopTimer()
}

func BenchmarkWorkerMem(b *testing.B) {
	w := NewWorker(number, b.N)
	ch := make(chan struct{})
	wg := runTest(b, func(i int, wg *sync.WaitGroup) {
		w.Push(newMemTask(ch, wg))
	})
	close(ch)
	w.Close()
	wg.Wait()
	b.StopTimer()
}

func BenchmarkGoroutineCPU(b *testing.B) {
	wg := runTest(b, func(i int, wg *sync.WaitGroup) {
		go func() {
			t := newTestTask(wg)
			t.Do()
		}()
	})
	wg.Wait()
	b.StopTimer()
}

func BenchmarkGoroutineMem(b *testing.B) {
	ch := make(chan struct{})
	wg := runTest(b, func(i int, wg *sync.WaitGroup) {
		go func() {
			t := newMemTask(ch, wg)
			t.Do()
		}()
	})
	close(ch)
	wg.Wait()
	b.StopTimer()
}

func BenchmarkDo(b *testing.B) {
	num := runtime.NumCPU()
	slice := make([]int, 0)
	ch := make(chan int, num)
	var wg sync.WaitGroup
	var mu sync.Mutex

	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()

	f := func() {
		defer wg.Done()
		defer mu.Unlock()
		mu.Lock()
		for i := range ch {
			slice = append(slice, i)
		}
	}

	wg.Add(num)
	for i := 0; i < num; i++ {
		go f()
	}

	wg.Wait()
	fmt.Println(len(slice))
}

func BenchmarkDo1(b *testing.B) {
	slice := make([]int, 0)
	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
	fmt.Println(len(slice))
}
