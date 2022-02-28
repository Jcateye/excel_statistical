package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/plandem/xlsx"
	"log"
	"os"
	"strconv"
	"time"
)

var personIndex, dateIndex, brandIndex, classIndex, productIndex, realShotIndex, linkIndex, orderIndex, receivableIndex, paidIndex, perforAccountIndex, paymentTypeIndex, payDateIndex, commentIndex = -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1

func main() {

	excel , excelHeadMap:= getSheetData()

	// Engin
	router := gin.Default()
	//router := gin.New()
	excelMap, _ := json.Marshal(excelHeadMap)

	log.Println(string(excelMap))
	router.Static("/web","./web")
	router.GET("/hello", func(context *gin.Context) {
		log.Println(">>>> hello gin start <<<<")
		context.JSON(200,gin.H{
			"code":200,
			"success":true,
			"data": excel,
			"head": excelHeadMap,
		})
	})
	// 指定地址和端口号
	router.Run("localhost:9090")

	//data, err := ioutil.ReadFile("config.txt")
	//if err != nil {
	//	fmt.Println("File reading error", err)
	//}
	//var dat map[string]string
	//if err := json.Unmarshal([]byte(data), &dat); err == nil {
	//	fmt.Println("==============json str 转map=======================")
	//	fmt.Println(dat["开始时间"])
	//	fmt.Println(dat["结束时间"])
	//}
	//
	//startTime, _ := time.Parse("2006-01-02", dat["开始时间"])
	//endTime, _ := time.Parse("2006-01-02", dat["结束时间"])
	//fmt.Println(startTime, endTime)
	//
	//file, err := xlsx.Open("市场部对账单.xlsx")
	//if err != nil {
	//	return
	//}
	//sheetsIter := file.Sheets()
	//perList := make([]Performance, 0)
	//for sheetsIter.HasNext() {
	//	sheetIndex, _ := sheetsIter.Next()
	//	sheet := file.Sheet(sheetIndex)
	//	//if sheet.Name() == "11月" {
	//	rowIter := sheet.Rows()
	//	for rowIter.HasNext() {
	//		rowIndex, _ := rowIter.Next()
	//		row := sheet.Row(rowIndex)
	//		cellIter := row.Cells()
	//		performance := Performance{}
	//		for cellIter.HasNext() {
	//			ci, _, _ := cellIter.Next()
	//			cell := row.Cell(ci)
	//			val := cell.Value()
	//			if val != "" {
	//				if rowIndex == 0 {
	//					// 初始化列index
	//					indexInit(ci, val)
	//				} else {
	//					rowToPerformance(&performance, ci, cell)
	//				}
	//				//fmt.Sprintf("row %v column %v value:%v", ri, ci, cell)
	//			}
	//		}
	//		//fmt.Println(performance)
	//		perList = append(perList, performance)
	//	}
	//	//}
	//	fmt.Printf("sheet.Name  %s %d \n", sheet.Name(), len(perList))
	//}
	//format := "2006年01月02日"
	//str := fmt.Sprintf("市场部\n%v-%v收到  \r\n", startTime.Format(format), endTime.Format(format))
	//statisMap := sheetStatical(perList, startTime, endTime)
	//receivable := 0.00
	//paid := 00.00
	//perforAccount := 00.00
	//realShot := 0
	//order := 0
	//for _, statis := range statisMap {
	//	//fmt.Printf(key)
	//	receivable += statis.receivable
	//	paid += statis.paid
	//	perforAccount += statis.perforAccount
	//	realShot += statis.realShot
	//	order += statis.order
	//	str += fmt.Sprintf("%v 应收金额：%v 已收金额：%v 实收业绩: %v \r\n", statis.person, statis.receivable, statis.paid, statis.perforAccount)
	//}
	//str += fmt.Sprintf("应收金额：%v 已收金额：%v 实收业绩: %v \r\n", receivable, paid, perforAccount)
	//str += fmt.Sprintf("收到订单：%v 实拍：%v  \r\n", realShot, order)
	//str += "------------------------------------------------\r\n"
	//WriteFile("result.txt", []byte(str))
	//file.Close()

}
func WriteFile(name string, data []byte) error {
	f, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err1 := f.Close(); err1 != nil && err == nil {
		err = err1
	}
	return err
}
func sheetStatical(performanceList []Performance, startTime time.Time, endTime time.Time) map[string]*Statistical {
	statisMap := make(map[string]*Statistical)
	for i := 0; i < len(performanceList); i++ {
		performance := performanceList[i]
		personName := performance.person
		if len(personName) == 0 {
			continue
		}
		date := performance.date
		if date.Equal(startTime) || date.Equal(endTime) || (date.After(startTime) && date.Before(endTime)) {
			_, ok := statisMap[personName]
			if !ok {
				statistical := Statistical{
					person: personName,
				}
				statisMap[personName] = &statistical
			}
			statistical := statisMap[personName]
			// 应收金额累计
			statistical.receivable += performance.receivable
			statistical.paid += performance.paid
			statistical.perforAccount += performance.perforAccount
			statistical.realShot += performance.realShot
			statistical.order += performance.order
		}
	}
	return statisMap
}

func rowToPerformance(performance *Performance, colIndex int, value *xlsx.Cell) {
	if colIndex == personIndex {
		performance.person = value.Value()
	} else if colIndex == dateIndex {
		performance.date, _ = value.Date()
	} else if colIndex == brandIndex {
		performance.brand = value.Value()
	} else if colIndex == productIndex {
		performance.product = value.Value()
	} else if colIndex == realShotIndex {
		performance.realShot, _ = strconv.Atoi(value.Value())
	} else if colIndex == linkIndex {
		performance.link, _ = strconv.Atoi(value.Value())
	} else if colIndex == orderIndex {
		performance.order, _ = strconv.Atoi(value.Value())
	} else if colIndex == receivableIndex {
		performance.receivable, _ = strconv.ParseFloat(value.Value(), 64)
	} else if colIndex == paidIndex {
		performance.paid, _ = strconv.ParseFloat(value.Value(), 64)
	} else if colIndex == perforAccountIndex {
		performance.perforAccount, _ = strconv.ParseFloat(value.Value(), 64)
	} else if colIndex == paymentTypeIndex {
		performance.paymentType = value.Value()
	} else if colIndex == commentIndex {
		performance.comment = value.Value()
	} else if colIndex == payDateIndex {
		performance.payDate, _ = value.Date()
		//fmt.Println("payDate date is ", performance.payDate)
	}
}

func indexInit(colIndex int, val string) {
	if val == "日期" {
		dateIndex = colIndex
	} else if val == "商标名称" {
		brandIndex = colIndex
	} else if val == "类别" {
		classIndex = colIndex
	} else if val == "产品" {
		productIndex = colIndex
	} else if val == "实拍" {
		realShotIndex = colIndex
	} else if val == "链接" {
		linkIndex = colIndex
	} else if val == "订单" {
		orderIndex = colIndex
	} else if val == "应收金额" {
		receivableIndex = colIndex
	} else if val == "实收金额" {
		paidIndex = colIndex
	} else if val == "销售业绩" {
		perforAccountIndex = colIndex
	} else if val == "付款方式" {
		paymentTypeIndex = colIndex
	} else if val == "付款日期" {
		payDateIndex = colIndex
	} else if val == "销售顾问" {
		personIndex = colIndex
	} else if val == "备注" {
		commentIndex = colIndex
	}

}
