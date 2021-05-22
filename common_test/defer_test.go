package common

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	panic("4")
}

func TestDefer1(t *testing.T) {
	i := "1"
	defer func() {
		fmt.Println(i)
	}()
	defer func() {
		i = i + "2"
		fmt.Println(i)
	}()
	defer func() {
		i = i + "3"
		fmt.Println(i)
	}()

	fmt.Println(i)
}

func TestDefer3(t *testing.T) {
	defer fmt.Println("1")
	defer func() {
		panic("5")
	}()
	defer fmt.Println("3")
	panic("4")
}
