package file_test

import (
	"fmt"
	"os"
	"testing"
)

func TestFileExist(t *testing.T) {
	_, err := os.Stat("ExcelFileDir") // os.Stat获取文件信息
	if err != nil {
		if !os.IsExist(err) {
			err := os.Mkdir("ExcelFileDir", os.ModePerm)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
