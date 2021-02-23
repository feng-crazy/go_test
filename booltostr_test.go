package go_test


import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestBoolToStr(t *testing.T){
	s := strconv.FormatBool(true)
	fmt.Println(s)
}


func TestInterface(t *testing.T) {
	var i interface{}
	i =  errors.New("hahah ")
	fmt.Printf("%+v", i)
}