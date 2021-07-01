package common

import (
	"fmt"
	"testing"
)

type myint int32
type myint32 = int32

func TestTypeAlias(t *testing.T) {
	var x int32 = 32.0
	var y = x
	// var y myint = x
	var z = x
	fmt.Println(y, z)
	var a myint = 10
	a = (myint)(x)
	a.Print()
}

func (i myint) Print() {
	fmt.Println(i)
}

// func (i myint32)Print(){
//
// }
