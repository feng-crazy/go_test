package main

/*
#include <stdlib.h>       // for free()
#include <stdint.h>

struct s2 {
	int32_t	 af;
};

struct s1 {
	int32_t	 a;
	int32_t	 b;
	struct s2 c;
};

// simple wrapper, see below
struct s1 call_c_struct(struct s1 s) {
	s.a *= 2;
	return s;
}
*/
import "C"

import (
	"fmt"
	// "unsafe"
)

type Sg2 struct {
	af int32
}

type Sg1 struct {
	a int32
	b int32
	c Sg2
}

type Sgc1 = C.struct_s1

func main() {

	s1 := C.struct_s1{a: 5} // or cast with C.int: s1 := C.struct_s1{C.int(i)}
	s1Ret := C.call_c_struct(s1)
	fmt.Printf("f31: s1=%v, s1Ret=%v\n", s1, s1Ret)
	fmt.Println(s1Ret.a, s1Ret.b, s1Ret.c)

	// var sg1Ret Sg1
	// // sg1 = Sg1{a: 5}
	// s2 := C.struct_s1{a: 4}
	// sg1Ret = C.call_c_struct(s2)
	// fmt.Printf("f31: s1=%v, s1Ret=%v\n", s2, sg1Ret)

	var Sgc1Ret Sgc1
	// sg1 = Sg1{a: 5}
	s3 := Sgc1{a: 4}
	Sgc1Ret = C.call_c_struct(s3)
	fmt.Printf("f31: s3=%v, Sgc1Ret=%v\n", s3, Sgc1Ret)
	fmt.Println(s1Ret.a, s1Ret.b, s1Ret.c)
}
