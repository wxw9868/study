package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
)

func main() {
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
