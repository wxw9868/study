package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println(math.MaxFloat32)
	rows := `2,Tina,37,"����,""����",Old Job
3,Alice Job,66,"""����Ӱ"",����","�Ϻ�,�Ϻ���"
4,John,44,"ϴ�»�101,""","LA""CITY"""
5,"Jane,li",55,Hiking,Canada`
	execute(rows)
}

func execute(rows string) {
	for _, row := range strings.Split(rows, "\n") {
		result := ""
		inPuote := false
		for i := 0; i < len(row); i++ {
			ch := string(row[i])
			if ch == "\"" {
				if inPuote {
					if i+1 < len(row) && string(row[i+1]) == "\"" {
						result += ch
						i++
					} else {
						inPuote = false
					}
				} else {
					inPuote = true
				}
			} else if ch == "," {
				if inPuote {
					result += ch
				} else {
					result += "\t"
				}
			} else {
				result += ch
			}
		}
		println(result)
	}
}
