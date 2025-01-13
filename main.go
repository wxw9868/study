package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
)

func main() {
	// // 假设你已经有了一个JSON格式的字符串
	// jsonString := `{"name":"John","age":30,"city":"New York"}`

	// var jsonBytes []byte = []byte(jsonString)
	// var formattedJSONBytes []byte

	// // 对JSON字符串进行格式化（缩进）
	// formattedJSONBytes, err := json.MarshalIndent(json.RawMessage(jsonBytes), "", "  ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// formattedJSONString := string(formattedJSONBytes)
	// fmt.Println(formattedJSONString)
	// return

	resp, err := http.Get("http://127.0.0.1:5678/api/word/cut?q=上海和深圳哪个城市幸福指数高")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Body)
	robots, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", robots)

	// 对JSON字符串进行格式化（缩进）
	formattedJSONBytes, err := json.MarshalIndent(json.RawMessage(robots), "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	formattedJSONString := string(formattedJSONBytes)
	fmt.Println(formattedJSONString)
	return

	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				fmt.Println(e)
				fmt.Println("mobile")
			case paint.Event:
				log.Print("Call OpenGL here.")
				a.Publish()
			}
		}
	})
}

func runPool() {
	var ch = make(chan bool)
	go write(ch)
	go read(ch)

	time.Sleep(2 * time.Second)
}

func write(ch chan bool) {
	ch <- true
}

func read(ch chan bool) {
	t := time.NewTimer(1 * time.Second)
	status := <-ch
	if status {
		//重新设置为5m
		if t.Reset(1 * time.Second) {
			fmt.Println("刷新成功。")
		}
	}
	go timerTask(t)
}

// 会话关闭定时器
func timerTask(timer *time.Timer) {
	for {
		select {
		case <-timer.C:
			fmt.Println("时间到")
		}
	}
}

func timer(roomId int, close bool) {
	t := time.NewTimer(5 * time.Second)
	if close {
		t.Reset(5 * time.Second)
	}
	for {
		select {
		case <-t.C:
			fmt.Println(roomId)
			fmt.Println("close")
		}
	}
}
