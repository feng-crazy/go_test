package go_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func TestExcel(t *testing.T){
	filename := "./data/devicemodel.xlsx"
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("model")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

}


func TestExcelWrite(t *testing.T){
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		println(err.Error())
	}
	styleID, err := file.NewStyle(`{"font":{"color":"#777777"}}`)
	if err != nil {
		println(err.Error())
	}
	if err := streamWriter.SetRow("A1", []interface{}{excelize.Cell{StyleID: styleID, Value: "Data"}, excelize.Cell{StyleID: styleID, Value: "Data1"}}); err != nil {
		println(err.Error())
	}
	for rowID := 2; rowID <= 1024; rowID++ {
		row := make([]interface{}, 50)
		for colID := 0; colID < 5; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, row); err != nil {
			println(err.Error())
		}
	}
	if err := streamWriter.Flush(); err != nil {
		println(err.Error())
	}
	if err := file.SaveAs("./data/Book1.xlsx"); err != nil {
		println(err.Error())
	}
}


func TestStringToI(t *testing.T){
	 fmt.Println(strconv.Atoi("A"))

	fmt.Println(strconv.ParseInt("A", 10, 0))

	str := "A"
	strRune := []rune(str)
	fmt.Println(strRune[0])

	colNum := int(strRune[0]) - 64
	fmt.Println(colNum)

	tmp, err := strconv.Atoi("")
	fmt.Println(tmp, err)
	if err != nil{
		eStr := err.Error()
		if eStr == `strconv.Atoi: parsing "": invalid syntax` {
			fmt.Println("haha")
		}
	}


	tmp1 , err := strconv.ParseFloat("10", 10)
	fmt.Println(tmp1, err)

	fmt.Println(strconv.ParseFloat("", 10))

	fmt.Println(strconv.ParseBool("True"))

	tagsMap := make(map[string]string)
	err = json.Unmarshal([]byte("{}"), &tagsMap)
	fmt.Println(tagsMap, err)
}

var ModelCellMap = map[string]string{
	"Name": "A",
	"Description": "B",
	"Tags": "C",
	"PropertyName": "D",
	"PropertyDescription": "E",
	"Type": "F",
	"AccessMode": "G",
	"DefaultValue": "H",
	"Minimum": "I",
	"Maximum": "J",
	"Unit": "K",
}

func TestExcel1(t *testing.T){
	var err error
	filename := "./data/test.xlsx"

	outputf := excelize.NewFile()

	index := outputf.NewSheet("test")

	outputf.SetActiveSheet(index)

	// streamWriter, err := outputf.NewStreamWriter("test")
	// if err != nil {
	//	return
	// }
	//
	// values := make([]interface{}, 1)
	// for k, v :=  range ModelCellMap {
	//	values[0] = k
	//	axis := v + strconv.Itoa(1)
	//	_ = streamWriter.SetRow(axis, values)
	//	fmt.Printf("axis:%s, value:%+v\n", axis, values)
	// }
	//
	// err = streamWriter.Flush()
	// if err != nil{
	//	fmt.Println(err)
	// }

	values := make([]interface{}, 1)
	for k, v :=  range ModelCellMap {
		values[0] = k
		axis := v + strconv.Itoa(1)
		_ = outputf.SetSheetRow("test", axis, &values)
		fmt.Printf("axis:%s, value:%+v\n", axis, values)
	}


	err = outputf.SaveAs(filename)
	if err != nil{
		fmt.Println(err)
	}

	inputf, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取 test 上所有单元格
	rows, err := inputf.GetRows("test")
	fmt.Println(rows[0])

	fmt.Println(inputf.GetCellValue("test", "A1"))

	fmt.Println(inputf.GetCellValue("test", "C3"))


}

func TestExcel2(t *testing.T) {
	var err error
	filename := "./data/device.xlsx"


	inputf, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println(inputf.GetCellValue("protocol", "A3"))


}