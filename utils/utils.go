package utils

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/thedevsaddam/gojsonq/v2"
)

func ParsingJSON() {
	// 书名、作者、分类
	var resUrl = "https://opendata.baidu.com/api.php?tn=wisexmlnew&dsp=iphone&alr=1&resource_id=5391&query="
	var querys = `赵旭李晴晴小说完整版免费阅读
赵旭李晴晴最新小说`
	var result []map[string]any
	var queryOther []string
	var queryArr = strings.Split(querys, "\n")
	for _, query := range queryArr {
		resp, err := http.Get(resUrl + url.QueryEscape(query))
		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		defer resp.Body.Close()

		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("err: ", err)
			return
		}

		jq := gojsonq.New().JSONString(string(bytes))
		if jq.Find("Result.[0]") != nil {
			jq = gojsonq.New().JSONString(string(bytes))
			bookname := jq.Find("Result.[0].DisplayData.resultData.tplData.content.bookinfo.[0].bookname")
			jq = gojsonq.New().JSONString(string(bytes))
			penname := jq.Find("Result.[0].DisplayData.resultData.tplData.content.bookinfo.[0].penname")
			jq = gojsonq.New().JSONString(string(bytes))
			category_raw := jq.Find("Result.[0].DisplayData.resultData.tplData.content.bookinfo.[0].category_raw")
			jq = gojsonq.New().JSONString(string(bytes))
			bd_tag := jq.Find("Result.[0].DisplayData.resultData.tplData.content.bookinfo.[0].bd_tag")

			categoryArr := strings.Split(fmt.Sprintf("%s", category_raw), "_")
			tagArr := strings.Split(fmt.Sprintf("%s", bd_tag), ",")
			category := categoryArr[1] + " " + tagArr[0]
			if len(tagArr) > 1 {
				category += " " + tagArr[1]
			}

			data := map[string]any{
				"query":    query,
				"bookname": bookname,
				"penname":  penname,
				"category": category,
			}
			result = append(result, data)
		} else {
			queryOther = append(queryOther, query)
		}
	}

	b, _ := json.Marshal(result)
	os.WriteFile("./data.json", b, 0666)
	o, _ := json.Marshal(queryOther)
	os.WriteFile("./queryOther_data.json", o, 0666)
}

func ReadWriteFile() {
	file, err := os.Open("query.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buffer.WriteString("'")
		buffer.WriteString(scanner.Text())
		buffer.WriteString("'")
		buffer.WriteString(",")
		// fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	if err = os.WriteFile("result.txt", buffer.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

// MethodRuntime 查看方法运行时长
func MethodRuntime(f func()) {
	start := time.Now()

	f()

	end := time.Since(start)

	fmt.Println(end)
}
