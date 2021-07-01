package common

import (
	"fmt"
	"testing"
)

func TestSliceAdd(t *testing.T) {
	// arrys := make([]string, 0)
	var arrys []string
	arrys = append(arrys, "test")
	fmt.Println(arrys)
}

func TestSliceSplit(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:]
	fmt.Println(b)

	b = a[1:2]
	fmt.Println(b)

	b = a[:4]
	fmt.Println(b)

	b = a[1:2:4]
	fmt.Println(b)
}
