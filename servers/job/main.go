package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/wxw9868/study/servers/job/database"
	"github.com/wxw9868/study/servers/job/excel"
)

func main() {
	WriteRegion()
	WriteUniversity()
}

type Region struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	PostCode string `json:"post_code"`
	Citys    []City `json:"citys"`
}

type City struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	PostCode string `json:"post_code"`
	Areas    []Area `json:"areas"`
}

type Area struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	PostCode string `json:"post_code"`
	Towns    []Town `json:"towns"`
}

type Town struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	PostCode string `json:"post_code"`
}

func WriteRegion() {
	f, err := os.Open("assets/ChinaCitys2025.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()

	var data []Region
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		fmt.Println(err)
		return
	}

	var i uint
	var regions []database.Region
	for _, region := range data {
		// fmt.Printf("region: %+v\n", region)
		v := region
		child := database.Region{
			Name:     v.Name,
			Code:     v.Code,
			PostCode: v.PostCode,
			Level:    1,
		}
		i = i + 1
		child.ID = i
		regions = append(regions, child)
		for _, city := range region.Citys {
			v1 := city
			child1 := database.Region{
				Name:     v1.Name,
				Code:     v1.Code,
				PostCode: v1.PostCode,
				ParentID: child.ID,
				Level:    2,
			}
			i = i + 1
			child1.ID = i
			regions = append(regions, child1)
			for _, area := range city.Areas {
				v2 := area
				child2 := database.Region{
					Name:     v2.Name,
					Code:     v2.Code,
					PostCode: v2.PostCode,
					ParentID: child1.ID,
					Level:    3,
				}
				i = i + 1
				child2.ID = i
				regions = append(regions, child2)
				for _, town := range area.Towns {
					v3 := town
					child3 := database.Region{
						Name:     v3.Name,
						Code:     v3.Code,
						PostCode: v3.PostCode,
						ParentID: child2.ID,
						Level:    4,
					}
					i = i + 1
					child3.ID = i
					regions = append(regions, child3)
				}
			}
		}
	}
	// fmt.Printf("regions: %d %+v\n", len(regions), regions[0])
	defer database.Close()
	if err := database.DB().CreateInBatches(&regions, 1000).Error; err != nil {
		log.Fatal("err: ", err)
	}

	if err := WriteFile(&regions, "regions.json"); err != nil {
		log.Fatal(err)
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

func WriteUniversity() {
	rows, err := excel.ReadExcelFile("assets/school.xlsx", "sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	var universities []University
	for _, row := range rows {
		num, _ := strconv.Atoi(row[0])
		if len(row) >= 6 && num >= 1 {
			// fmt.Println(row)
			universitie := University{
				SchoolName:          row[1],
				SchoolIdentifier:    row[2],
				CompetentDepartment: row[3],
				Location:            row[4],
				SchoolLevel:         row[5],
			}
			if len(row) == 7 {
				universitie.Note = row[6]
			}
			universities = append(universities, universitie)
		}
	}

	if err := database.DB().Create(&universities).Error; err != nil {
		log.Fatal("err: ", err)
	}
	if err := WriteFile(&universities, "universities.json"); err != nil {
		log.Fatal(err)
	}
}

func WriteFile(v any, name string) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if err := os.WriteFile(name, bytes, 0644); err != nil {
		return err
	}
	return nil
}
