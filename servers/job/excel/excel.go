package excel

import (
	"errors"

	"github.com/xuri/excelize/v2"
)

func ReadExcelFile(filename, sheet string) ([][]string, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, errors.New("open excel file error")
	}
	defer func() {
		if err = f.Close(); err != nil {
			panic(err)
		}
	}()

	// 获取 sheet 上所有单元格
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, errors.New("get rows error")
	}
	return rows, nil
}

func WriteExcelFile(filename, sheet string, data [][]string) error {
	return nil
}
