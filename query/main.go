package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/thedevsaddam/gojsonq/v2"
)

// 书名、作者、分类
var resUrl = "https://opendata.baidu.com/api.php?tn=wisexmlnew&dsp=iphone&alr=1&resource_id=5391&query="
var querys = `赵旭李晴晴小说完整版免费阅读
赵旭李晴晴最新小说
太荒吞天诀最新全文
顾北弦苏婳免费阅读
太荒吞天诀免费观看
苏婳顾北弦结局
苏婳顾北弦免费阅读
苏婳顾北弦免费阅读全文苏婳反转
《太荒吞天诀》免费阅读
吞天神鼎小说免费阅读
洛尘最新章节免费阅读
吞天神鼎最新章节
太荒吞天诀全文阅读
赵旭李晴晴
超级女婿赵旭李晴晴小说完整版免费阅读
太荒吞天诀最后结局
《离婚后她惊艳了世界》小说
吞天神鼎最新章节更新
赵旭李晴晴小说免费阅读,全集
李晴晴赵旭的小说免费阅读
龙王医婿最新章节
离婚后她惊艳了世界 明婳
古小暖江尘御最新章节
太荒吞天诀结局
楚浩苏念全文免费阅读
赵旭李晴晴最新章节
顾北弦苏婳免费阅读无弹窗
离婚后她惊艳了世界免费阅读
棺香美人最新章节
苏奕苏玄钧免费阅读最新
顾北弦苏婳小说免费阅读无弹窗
《上门龙婿》免费阅读
宁天林冉冉全文免费
全能兵王萧晨免费阅读
楚皓苏念小说免费阅读最新章节
赵旭李晴晴最新小说免费阅读
吞天神鼎
离婚后她惊艳了世界小说免费阅读全文
苏熙凌久泽全文免费阅读
道界天下最新章节
卢丹妮邓佳哲免费阅读
柳无邪徐凌雪吞天神鼎全文免费阅读
战神狂飙全集完整版
苏熙凌久泽小说免费阅读
苏奕文灵雪小说免费阅读
柳无邪徐凌雪小说全文免费阅读
唐楚楚的小说推荐
都市仙尊洛尘最新章节
武神主宰最新章节列表
盖世神医小说免费阅`

var result []map[string]interface{}
var queryOther []string

func main() {
	queryArr := strings.Split(querys, "\n")
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

			data := map[string]interface{}{
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
