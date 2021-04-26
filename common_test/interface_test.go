package common

import (
	"fmt"
	"testing"
)

func TestInterfaceValue(t *testing.T) {
	b := true
	var i interface{}
	i = b
	if i == true {
		fmt.Println("1111111111")
	} else {
		fmt.Println("0000000000")
	}
}

func TestInterfaceValue1(t *testing.T) {
	s := "true"
	var i interface{}
	i = s
	if i == "true" {
		fmt.Println("1111111111")
	} else {
		fmt.Println("0000000000")
	}
}

func TestInterfaceValue2(t *testing.T) {
	var s int
	s = 2
	var i interface{}
	i = s
	var b = 2
	if i != b {
		fmt.Println("1111111111")
	} else {
		fmt.Println("0000000000")
	}
}
