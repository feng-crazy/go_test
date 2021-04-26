package excel_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func TestExcel(t *testing.T) {
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

func TestExcelWrite(t *testing.T) {
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

func TestStringToI(t *testing.T) {
	fmt.Println(strconv.Atoi("A"))

	fmt.Println(strconv.ParseInt("A", 10, 0))

	str := "A"
	strRune := []rune(str)
	fmt.Println(strRune[0])

	colNum := int(strRune[0]) - 64
	fmt.Println(colNum)

	tmp, err := strconv.Atoi("")
	fmt.Println(tmp, err)
	if err != nil {
		eStr := err.Error()
		if eStr == `strconv.Atoi: parsing "": invalid syntax` {
			fmt.Println("haha")
		}
	}

	tmp1, err := strconv.ParseFloat("10", 10)
	fmt.Println(tmp1, err)

	fmt.Println(strconv.ParseFloat("", 10))

	fmt.Println(strconv.ParseBool("True"))

	tagsMap := make(map[string]string)
	err = json.Unmarshal([]byte("{}"), &tagsMap)
	fmt.Println(tagsMap, err)
}

var ModelCellMap = map[string]string{
	"Name":                "A",
	"Description":         "B",
	"Tags":                "C",
	"PropertyName":        "D",
	"PropertyDescription": "E",
	"Type":                "F",
	"AccessMode":          "G",
	"DefaultValue":        "H",
	"Minimum":             "I",
	"Maximum":             "J",
	"Unit":                "K",
}

func TestExcel1(t *testing.T) {
	var err error
	filename := "./data/test.xlsx"

	outputf := excelize.NewFile()

	index := outputf.NewSheet("test")

	outputf.SetActiveSheet(index)

	values := make([]interface{}, 1)
	for k, v := range ModelCellMap {
		values[0] = k
		axis := v + strconv.Itoa(1)
		_ = outputf.SetSheetRow("test", axis, &values)
		fmt.Printf("axis:%s, value:%+v\n", axis, values)
	}
	styleID, _ := outputf.NewStyle(`{"alignment":{"horizontal":"center","Vertical":"center"},"font":{"bold":true}}`)
	_ = outputf.SetCellStyle("test", "A1", "E1", styleID)

	v1, _ := outputf.GetCellValue("test", "B1")
	fmt.Println("hahahhahah v1: ", v1)

	_ = outputf.SetSheetRow("test", "A2", &[]interface{}{"hehe a"})
	_ = outputf.SetSheetRow("test", "B2", &[]interface{}{"hehe b"})
	_ = outputf.SetSheetRow("test", "C2", &[]interface{}{"hehe c"})
	_ = outputf.SetSheetRow("test", "D2", &[]interface{}{"hehe d"})
	_ = outputf.SetSheetRow("test", "E2", &[]interface{}{"hehe e"})

	_ = outputf.MergeCell("test", "A2", "E2")

	dvRange := excelize.NewDataValidation(false)
	dvRange.Sqref = "A5:A1048576"
	err = dvRange.SetDropList([]string{"1", "2", "3"})
	err = outputf.AddDataValidation("test", dvRange)
	err = outputf.AddComment("test", "A3", `{"author":"Excelize: ","text":"This is a comment."}`)

	v, _ := outputf.GetCellValue("test", "C2")
	fmt.Println("hahahhahah v: ", v)
	err = outputf.SaveAs(filename)
	if err != nil {
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

	rows, _ := inputf.Rows("")
	rows.Columns()
	fmt.Println(inputf.GetCellValue("protocol", "A3"))

}

func TestXinNeng(t *testing.T) {
	var err error
	filename := "./data/test.xlsx"

	outputf := excelize.NewFile()

	index := outputf.NewSheet("test")

	outputf.SetActiveSheet(index)

	t1 := time.Now()

	for i := 1; i < 100; i++ {
		// axis, _ := excelize.CoordinatesToCellName(1, i)
		// axis1, _ := excelize.CoordinatesToCellName(2, i)
		// _ = outputf.MergeCell("test", axis, axis1)
		for j := 1; j < 50; j++ {
			axis, _ := excelize.CoordinatesToCellName(j, i)
			_ = outputf.SetCellValue("test", axis, i*j)
			// styleID, err := file.NewStyle(`{"font":{"color":"#777777"}}`)
			styleID, _ := outputf.NewStyle(`{"alignment":{"horizontal":"center","Vertical":"center"},"font":{"color":"#A9A9A9","bold":true}}`)
			_ = outputf.SetCellStyle("test", axis, axis, styleID)
		}
		// 合并单元格

	}

	t2 := time.Now()
	fmt.Println(t2.Sub(t1))

	err = outputf.SaveAs(filename)
	if err != nil {
		fmt.Println(err)
	}
}

func TestReadXinNeng(t *testing.T) {
	var err error
	filename := "./data/modbus_3int-modbus_temp (2).xlsx"

	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	// sheetName := "test"
	sheetName := "device"
	allrows, _ := f.GetRows(sheetName)
	fmt.Println(len(allrows))

	rows, err := f.Rows(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			fmt.Println(err)
		}
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

}
