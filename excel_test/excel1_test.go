package excel_test

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/sirupsen/logrus"
)

type Excel struct {
	FileName  string
	SheetName string
	Execfile  *excelize.File
}

func NewExcelMarshal(fileName string, sheetName string) (e Excel) {
	e.FileName = fileName
	e.SheetName = sheetName
	e.Execfile = excelize.NewFile()

	index := e.Execfile.NewSheet(sheetName)

	e.Execfile.SetActiveSheet(index)
	return
}

func NewExcelParse(fileName string, sheetName string) (e Excel, err error) {
	e.FileName = fileName
	e.SheetName = sheetName
	e.Execfile, err = excelize.OpenFile(fileName)

	index := e.Execfile.NewSheet(sheetName)

	e.Execfile.SetActiveSheet(index)
	return
}

func (e *Excel) SetSheetValue(axis string, value interface{}) {
	err := e.Execfile.SetSheetRow(e.SheetName, axis, &[]interface{}{value})
	if err != nil {
		logrus.Error(err)
	}
}

func (e *Excel) GetDeviceSheetValue(axis string) string {
	value, err := e.Execfile.GetCellValue(e.SheetName, axis)
	if err != nil {
		logrus.Error(err)
	}
	return value
}

func (e *Excel) FileSave() error {
	if err := e.Execfile.SaveAs(e.FileName); err != nil {
		logrus.Error(err.Error())
		return err
	}
	return nil
}
