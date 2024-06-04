package goquery

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
)

type Data struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	School   string `json:"school"`
	Level    string `json:"level"`
	Stype    string `json:"stype"`
	Ways     string `json:"ways"`
	Duration string `json:"duration"`
	Number   string `json:"number"`
}

func GoQuery() {
	// Request the HTML page.
	res, err := http.Get("http://58.129.246.60:8888/web/specifyReportedInfo/specify_reported_info.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	table := doc.Find("#AutoNumber4")

	var datas []Data
	// Find the review items
	table.Find("tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		var data Data
		s.Find("td").Each(func(i int, c *goquery.Selection) {
			if i == 0 {
				data.Code = c.Text()
			}
			if i == 1 {
				data.Name = c.Text()
			}
			if i == 2 {
				data.School = c.Text()
			}
			if i == 3 {
				data.Level = c.Text()
			}
			if i == 4 {
				data.Stype = c.Text()
			}
			if i == 5 {
				data.Ways = c.Text()
			}
			if i == 6 {
				data.Duration = c.Text()
			}
			if i == 7 {
				data.Number = c.Text()
			}
		})

		datas = append(datas, data)
	})

	b, err := json.Marshal(datas)
	if err = os.WriteFile("demo.json", b, 0666); err != nil {
		log.Fatal(err)
	}
}
