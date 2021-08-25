package common

import (
	"fmt"
	"testing"
)

type TP struct {
	value int
	key   string
}

func ModifyTp(p *TP) {
	p.value = 100
	p.key = "100"
}

func TestPoint(t *testing.T) {
	p := TP{
		value: 0,
		key:   "0",
	}

	ModifyTp(&p)

	fmt.Println(p)
}
