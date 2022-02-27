package main

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/plandem/xlsx"
	"strconv"
	"strings"
	"time"
)

func getSheetData() (excel map[string][][]string, excelHeadMap map[string]map[int]TableHead) {
	file, err := xlsx.Open("市场部对账单.xlsx")
	if err != nil {
		return excel, excelHeadMap
	}
	excel = make(map[string][][]string)
	excelHeadMap = make(map[string]map[int]TableHead)
	sheetsIter := file.Sheets()
	for sheetsIter.HasNext() {
		sheetArr := make([][]string, 0)
		sheetIndex, _ := sheetsIter.Next()
		sheet := file.Sheet(sheetIndex)
		rowIter := sheet.Rows()
		headMap := make(map[int]TableHead)
		head := true
		for rowIter.HasNext() {
			rowValueArr := make([]string, 0)

			var headInfo TableHead
			rowIndex, _ := rowIter.Next()
			row := sheet.Row(rowIndex)
			cellIter := row.Cells()
			index := 0
			for cellIter.HasNext() {
				ci, _, _ := cellIter.Next()
				cell := row.Cell(ci)
				val := cell.Value()
				// 第一次保存下表头/
				if head {
					// 每个表头格内容
					set := mapset.NewSet()
					headInfo = TableHead{
						index:    index,
						name:     val,
						ValueSet: set}
					headMap[index] = headInfo
				} else {
					if headInfo, ok := headMap[index]; ok {
						if strings.Contains(headInfo.name, "日期") {
							if val != "" {
								valDate := excelDateToDate(val)
								val = valDate.Format("2006-01-02")
							}
						}
						headInfo.ValueSet.Add(val)
					}
				}
				index++
				rowValueArr = append(rowValueArr, val)
				//fmt.Printf("type is ----- %s , %v \n", reflect.TypeOf(val), val)
			}
			head = false
			//fmt.Println(performance)
			// 每一行数据放入表map
			sheetArr = append(sheetArr, rowValueArr)
		}
		fmt.Printf("%s  %+v \n", sheet.Name(), headMap)
		excelHeadMap[sheet.Name()] = headMap

		//}
		//fmt.Printf("sheet.Name  %s %d \n", sheet.Name(), len(perList))
		excel[sheet.Name()] = sheetArr
	}
	fmt.Println(excelHeadMap)
	return excel, excelHeadMap
}

func excelDateToDate(excelDate string) time.Time {
	excelTime := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	var days, _ = strconv.Atoi(excelDate)
	return excelTime.Add(time.Second * time.Duration(days*86400))
}
