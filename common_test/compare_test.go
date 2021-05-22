package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInterfaceCompare(t *testing.T) {
	var i1, i2 interface{}
	i1 = 10
	// var f2 float64 = 10.0

	i2 = "test"
	tmp, ok := i2.(float64)
	var rhsf int
	rhsf = int(tmp)

	fmt.Println(rhsf, ok)

	res := reflect.DeepEqual(i1, i2)

	fmt.Println(res)
}
