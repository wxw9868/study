package utils

import (
	"fmt"
	"time"
)

// MethodRuntime 查看方法运行时长
func MethodRuntime(f func()) {
	start := time.Now()

	f()

	end := time.Since(start)

	fmt.Println(end)
}
