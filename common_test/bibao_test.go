package common

import (
	"fmt"
	"testing"
)

func app() func(string) string {
	t := "Hi"
	c := func(b string) string {
		fmt.Printf("%v, %s\n", &t, t)
		t = t + " " + b
		return t
	}
	return c
}

func TestBiBao(t *testing.T) {
	a := app()
	b := app()
	a("go")
	fmt.Println(b("All"))
}

func hello(i *int) int {
	defer func() {
		*i = 19
	}()
	return *i
}

func hello1(i *int) (j int) {
	defer func() {
		j = 19
	}()
	j = *i
	return j
}

func hello2(i *int) (j int) {
	defer func() {
		j = 19
	}()
	return *i
}

func TestBibaoDefer(t *testing.T) {
	i := 10
	j := hello(&i)
	fmt.Println(i, j)

	i = 10
	j = hello1(&i)
	fmt.Println(i, j)

	i = 10
	j = hello2(&i)
	fmt.Println(i, j)
}
