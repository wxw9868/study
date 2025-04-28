package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("/mnt/c/Users/wxw9868/Downloads/school.xlsx")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 获取工作表中指定单元格的值
	// cell, err := f.GetCellValue("sheet1", "B3")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("cell: ", cell)

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	var universities []University

	for _, row := range rows {
		if len(row) == 6 {
			universities = append(universities, University{
				SchoolName:          row[1],
				SchoolIdentifier:    row[2],
				CompetentDepartment: row[3],
				Location:            row[4],
				SchoolLevel:         row[5],
			})
		}
		// for _, colCell := range row {
		// 	fmt.Print(colCell, "\t")
		// }
		// fmt.Println()
	}

	bytes, err := json.Marshal(&universities)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := os.WriteFile("universities.json", bytes, 0644); err != nil {
		fmt.Println(err)
		return
	}
}

type University struct {
	SchoolName          string `gorm:"column:school_name;type:varchar(100);uniqueIndex;not null;comment:学校名称"`
	SchoolIdentifier    string `gorm:"column:school_identifier;type:varchar(30);uniqueIndex;not null;comment:学校标识码"`
	CompetentDepartment string `gorm:"column:competent_department;type:varchar(100);index;comment:主管部门"` // Added index
	Location            string `gorm:"column:location;type:varchar(100);index;comment:所在地"`              // Added index
	SchoolLevel         string `gorm:"column:school_level;type:varchar(20);index;comment:办学层次"`          // Added index
	Note                string `gorm:"column:note;type:varchar(255);comment:备注"`
}
