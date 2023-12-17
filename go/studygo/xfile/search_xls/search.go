package search_xls

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
)

func SearchXls(fileName string) error {

	fmt.Println("searchXls start")
	file, err := excelize.OpenFile(fileName)
	if err != nil {
		return err
	}

	fmt.Println("searchXls mid")

	var thisSheet = sheet1

	cols, err := file.GetCols(thisSheet)

	rows, err := file.GetRows(thisSheet)
	if err != nil {
		return err
	}

	var mp = map[int]bool{}
	var locationMp = map[int]bool{}
	var finalMp = map[int]bool{}

	for _, col := range cols {

		if len(col) > 0 && col[1] == "专业" {
			for i, row := range col {
				if strings.Contains(row, "电子信息") || strings.Contains(row, "不限") {
					mp[i] = true
				}
			}
		}

		if len(col) > 0 && col[1] == "工作地点" {
			for i, row := range col {
				if strings.Contains(row, "广东省") {
					locationMp[i] = true
				}
			}
		}

	}

	fmt.Println("mp:", mp, "locationMp:", locationMp)

	for k, v := range mp {
		if v && locationMp[k] {
			finalMp[k] = true
		}
	}

	for k := range finalMp {
		fmt.Println(rows[k])
	}

	fmt.Println("row[1]", rows[1])

	f := excelize.NewFile()

	index, _ := f.NewSheet("Sheet1")

	err = f.SetSheetRow("Sheet1", "A1", &rows[1])
	if err != nil {
		return err
	}

	i := 2
	for k, _ := range finalMp {
		err = f.SetSheetRow("Sheet1", "A"+strconv.Itoa(i), &rows[k])
		i++
	}

	f.SetActiveSheet(index)

	f.SaveAs("/Users/zp/Documents/knowledgeTree/go/studygo/xfile/search_xls/" + thisSheet + ".xlsx")

	if err != nil {
		return err
	}

	return nil
}
