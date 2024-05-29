package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	MethodRuntime(TraverseFolders)
	// MethodRuntime(ReusingGoroutine)
	// MethodRuntime(UsingMutexLock)
}

// 遍历文件夹
func TraverseFolders() {
	path := "/Users/v_weixiongwei/go/src/video/assets/image/poster"
	recursive(path)
}

func recursive(dir string) {
	var data = make(map[string]struct{})
	var actress = make(map[string]struct{})

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			recursive(dir + "/" + file.Name())
		}
		name := file.Name()
		if filepath.Ext(name) == ".jpg" {
			arr := strings.Split(strings.Split(name, ".")[0], "_")
			actress[arr[len(arr)-1]] = struct{}{}
			if len(arr[len(arr)-2]) < 20 && len(arr[len(arr)-2]) > 6 {
				_, ok := actress[arr[len(arr)-2]]
				if !ok {
					data[arr[len(arr)-2]] = struct{}{}
					// fmt.Printf("%q, %q\n", arr[len(arr)-2], arr[len(arr)-1])
				}
			}
		}
	}
	WriteMapToFile("actress.json", actress)
	WriteMapToFile("data.json", data)

	// ReadFileToMap("actress.json", actress)
	// var slice []string
	// for k, _ := range actress {
	// 	slice = append(slice, k)
	// }
	// sort.Strings(slice)
	// fmt.Printf("%+v\n", slice)

	// for k, _ := range actress {
	// 	_, ok := actress[k]
	// 	if !ok {
	// 		actress[k] = struct{}{}
	// 	}
	// }
}

func ReadFileToMap(name string, v any) error {
	bytes, err := os.ReadFile(name)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bytes, &v); err != nil {
		return err
	}
	return nil
}

func WriteMapToFile(name string, v any) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	err = os.WriteFile(name, bytes, 0666)
	if err != nil {
		return err
	}
	return nil
}

// ReusingGoroutine 重复使用协程
func ReusingGoroutine() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var ch = make(chan int)
	var data = make(map[int]int, 0)
	// data := new(sync.Map)
	// data := make([]int, 0)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
		E:
			for {
				r, ok := <-ch
				if ok {
					// 使用互斥锁保护 data
					mu.Lock()
					data[r] = r
					// data.LoadOrStore(r, r)
					// data = append(data, r)
					mu.Unlock()
				} else {
					fmt.Println(i, r, ok)
					break E
				}
			}
		}(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Wait()

	fmt.Println(data)
}

// UsingMutexLock 使用互斥锁
func UsingMutexLock() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var data = make([]int, 0)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// 使用互斥锁保护 data
			mu.Lock()
			data = append(data, i)
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(data)
}

// 查看方法运行时长
func MethodRuntime(f func()) {
	start := time.Now()

	f()

	end := time.Since(start)

	fmt.Println(end)
}
