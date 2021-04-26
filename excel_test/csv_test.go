package excel_test

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCsv(t *testing.T) {
	fileName := "./data/modbus.csv"
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err.Error())
	}
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	row, _ := r2.Read()
	fmt.Println(row)
	fmt.Println("//////")
	ss, _ := r2.ReadAll()
	// fmt.Println(ss)
	sz := len(ss)
	for i := 0; i < sz; i++ {
		fmt.Println(ss[i])
	}
}
