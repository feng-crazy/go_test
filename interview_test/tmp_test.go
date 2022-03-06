package interview_test

import (
	"fmt"
	"testing"
)

func TestInterview1(t *testing.T) {
	const (
		x = iota
		_
		y
		z = "zz"
		k
		p = iota
	)

	fmt.Println(x, y, z, k, p)
}

func TestInterview2(t *testing.T) {
	type person struct {
		name string
	}

	var m map[person]int
	p := person{"make"}
	fmt.Println(m[p])
	//m[p] = 1

	a := [5]int{1, 2, 3, 4, 5}
	tb := a[3:4:4]
	fmt.Println(tb[0])

	a1 := []int{1, 2, 3, 4, 5}
	tb1 := a1[3:4:4]
	fmt.Println(tb1[0], cap(tb1))

	a2 := []int{1, 2, 3, 4, 5}
	tb2 := a2[3:4]
	fmt.Println(tb2[0], cap(tb2))
}

func TestInterview3(t *testing.T) {
	a := [2]int{5, 6}
	b := [3]int{5, 6}
	fmt.Println(a, b)
	//if a == b {
	//	fmt.Println("equal")
	//} else {
	//	fmt.Println("not equal")
	//}

	a1 := [2]int{5, 6}
	b1 := [2]int{5, 6}

	if a1 == b1 {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	a2 := make([]int, 2)
	fmt.Println(a2)
}

func TestInterview4(t *testing.T) {
	var c = make(chan int, 1)
	var i1 interface{} = c

	if i1 == c {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	var str string = "test"
	ss := []rune(str)
	ss[0] = 'x'
	fmt.Println(ss)
}

func TestInterview5(t *testing.T) {
	var i1 int = 1
	var p *int = &i1
	var i interface{} = p

	if p == i {
		fmt.Println("p == i")
	}
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

type Person struct {
	age int
}

func TestInterview6(t *testing.T) {

	//fmt.Println(triple(4)) // "12"

	person := Person{28}

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(&person)

	// 3.
	defer func() {
		fmt.Println(person)
		fmt.Println(person.age)
	}()

	person.age = 29
}

func TestInterview7(t *testing.T) {

	//fmt.Println(triple(4)) // "12"

	person := Person{28}

	f := func() {
		fmt.Println(person)
	}
	person.age = 29
	f()
}

func TestInterview8(t *testing.T) {

	//fmt.Println(triple(4)) // "12"

	person := 28

	f := func() {
		fmt.Println(person)
	}
	person = 29
	f()
}

func TestInterview9(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1)
	if a := 1; false {
	} else if b := 2; false {
	} else {
		println(a, b)
	}
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return []string{"North", "East", "South", "West"}[d]
}

func TestInterview10(t *testing.T) {
	fmt.Println(South)

	v := []int{1, 2, 3}
	for i, vv := range v {
		if i+1 < len(v) {
			v[i+1] = 100 + 1
		}
		fmt.Println("vv: ", vv)
	}
	fmt.Println("v", v)

	for i, vv := range v {
		v = append(v, i)
		fmt.Println(vv)
	}
	fmt.Println(v)

	//for i, vv := 1, 2.0; i< 2; i++ {
	//	fmt.Println(vv)
	//}
	//
	//i := 0
	//v := 22.2
	//for i := 1, vv := 2; i< 2; i++ {
	//	fmt.Println(vv)
	//}

	var m = map[string]int{
		"B": 22,
		"C": 23,
		"A": 21,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}

func change(s ...int) {
	fmt.Println("change f:", s, len(s), cap(s))
	s = append(s, 3)
	fmt.Println("change after:", s, len(s), cap(s))
}

func TestInterview11(t *testing.T) {
	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)

	//var b bool
	////b = bool(0)
	//
	//b = (1==2)
	//
	//fmt.Println(b)
	//
	//var i int
	//i = int(b)
}

func TestInterview12(t *testing.T) {
	//int_chan := make(chan int)
	//string_chan := make(chan string)

	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	fmt.Println(len(int_chan), cap(int_chan))
	//go func() {
	//	int_chan <- 1
	//	string_chan <- "hello"
	//	fmt.Println("hello")
	//}()

	int_chan <- 1
	string_chan <- "hello"

	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		fmt.Println(value)
	}
}

func readCh(c <-chan int) {
	//close(c)
	<-c
}

func write(c chan<- int) {
	close(c)
}
func TestInterview13(t *testing.T) {
	//var c chan int = nil
	//c <- 1
	//i := <- c
	//fmt.Println(i)

	c1 := make(chan int, 0)

	write(c1)
	i, ok := <-c1
	fmt.Println(i, ok)
}

func TestInterview14(t *testing.T) {
	type Param map[string]interface{}

	type Show struct {
		*Param
	}
	//s := new(Show)
	//s.Param["day"] = 2
	s := new(Show)
	// 修复代码
	p := make(Param)
	p["day"] = 2
	s.Param = &p
	//tmp := *s.Param
	//fmt.Println(s.Param["day"])
}
