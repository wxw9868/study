package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

var savePath = "avatar"

func main() {
	start := time.Now().Local()

	var wg = new(sync.WaitGroup)
	var urls = make(chan string)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go Work(urls, wg)
	}

	wg.Add(1)
	go BatchDownloadImages(urls, wg)

	wg.Wait()
	fmt.Printf("SUCCESS, elapsed: %v\n", time.Since(start))

}

func Work(urls chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		err := DownloadImage(url, savePath, "")
		if err != nil {
			fmt.Printf("Image download failed: %s, error: %s\n", url, err)
		}
	}
}

func BatchDownloadImages(urls chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var (
		items    uint32
		requests uint32
		success  uint32
		failure  uint32
		results  uint32
	)

	c := colly.NewCollector(colly.UserAgent(browser.Random()), colly.AllowURLRevisit())

	q, _ := queue.New(
		runtime.NumCPU(),
		&queue.InMemoryQueueStorage{MaxSize: 100000},
	)

	c.OnHTML(".sewhjer > img", func(e *colly.HTMLElement) {
		link := strings.Join([]string{e.Request.URL.Scheme, "://", e.Request.URL.Host, e.Attr("data-src")}, "")

		urls <- link
		atomic.AddUint32(&results, 1)
		// fmt.Printf("Link found: %s -> %s\n", link, filepath.Ext(dataSrc))
	})

	c.OnRequest(func(r *colly.Request) {
		atomic.AddUint32(&requests, 1)
	})
	c.OnResponse(func(resp *colly.Response) {
		if resp.StatusCode == http.StatusOK {
			atomic.AddUint32(&success, 1)
		} else {
			atomic.AddUint32(&failure, 1)
		}
	})
	c.OnError(func(resp *colly.Response, err error) {
		atomic.AddUint32(&failure, 1)
	})

	var url string
	for i := 1; i < 48; i++ {
		if i > 2 {
			url = fmt.Sprintf("https://www.gnmxjj.com/articlecolumn/starziliaoku_a%d.html", i)
		} else {
			url = "https://www.gnmxjj.com/articlecolumn/starziliaoku.html"
		}
		q.AddURL(url)
		atomic.AddUint32(&items, 1)
	}

	if err := q.Run(c); err != nil {
		log.Fatalf("Queue.Run() return an error: %v", err)
	}

	close(urls)
	fmt.Printf("wrong Queue implementation: items = %d, requests = %d, success = %d, failure = %d, results = %d\n", items, requests, success, failure, results)
}

func DownloadImage(url, savePath, saveFile string) error {
	if err := os.MkdirAll(savePath, 0750); err != nil {
		log.Fatal(err)
	}

	// 发起 GET 请求获取图片数据
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("wrong http status code: %d", resp.StatusCode)
	}

	if saveFile == "" {
		// 获取原文件名
		_, saveFile = path.Split(resp.Request.URL.Path)
	}

	// 创建保存图片的文件
	file, err := os.Create(path.Join(savePath, saveFile))
	if err != nil {
		return err
	}
	defer file.Close()

	// 将响应体的数据写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
