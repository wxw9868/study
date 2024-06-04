package redis

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// 测试
func BenchmarkRedisLock1(b *testing.B) {
	x := b.N
	wg := sync.WaitGroup{}
	wg.Add(x)
	now := time.Now()

	var n int64
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	for i := 0; i < x; i++ {
		go func() {
			defer wg.Done()
			rl := NewRedisLock(ctx, "testLock")
			if err := rl.Lock(); err != nil {
				cancel() // 这里cancel，可以在第一个超时或者发生错误后，后边的不再尝试加锁。
				return
			}
			defer rl.Unlock()
			atomic.AddInt64(&n, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("测试数:%d\t成功数:%d\t结果:%t\t耗时： %d \n", x, n, x == int(n), time.Since(now).Milliseconds())
}
