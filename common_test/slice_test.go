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

	b = a[0:5]
	fmt.Println(b)

	c := []int{}
	b = c[0:0]
	fmt.Println(b)
}

func TestSliceCopy(t *testing.T) {
	// arrys := make([]string, 0)
	var arrys = make([]string, 1)
	fmt.Printf("%p\n", arrys)
	fmt.Printf("%p\n", &arrys[0])

	arrys = append(arrys, "test")
	fmt.Printf("%p\n", arrys)
	fmt.Printf("%p\n", &arrys[0])

	var arr = arrys
	arrys = make([]string, 0)
	fmt.Println(arr[1])
}

func Copy(arr []string) {
	var array = make([]string, 0)
	array = append(array, "ttttt")
	arr = array
	arr[0] = "gggggg"
}

func TestSliceCopy1(t *testing.T) {
	var arrys = make([]string, 1)
	Copy(arrys)
	fmt.Println(arrys)

}

func TestSliceCopy2(t *testing.T) {
	var arrys = make([]string, 1)
	arrys[0] = "test"

	var arr = arrys
	arr[0] = "gggg"
	arr = make([]string, 0)
	fmt.Println(arrys)
}
