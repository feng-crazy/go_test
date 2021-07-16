package main

/*
#include "libtest.h"
#include <stdlib.h>       // for free()
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -ltest -lstdc++
*/
import "C"

func main() {}

//export number_add_mod
func number_add_mod(a, b, mod C.int) C.int {
	return C.number_add_mod(a, b, mod)
}

// 无法将C打包成库， go来调用，然后把go打包成库，C来调用
