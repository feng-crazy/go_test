package go_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestToStr(t *testing.T) {
	s, err := strconv.Atoi("")
	fmt.Println(s, err)
}

func TestBoolToStr(t *testing.T) {
	s := strconv.FormatBool(true)
	fmt.Println(s)
}

func TestInterface(t *testing.T) {
	var i interface{}
	i = errors.New("hahah ")
	fmt.Printf("%+v", i)
}

func TestTmp(t *testing.T) {
	data := 1
	input := data

	{
		var err error
		input, err = 2, errors.New("test")
		fmt.Println(err)
		fmt.Println(input)
	}

	fmt.Println(input)
}
