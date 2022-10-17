package channel

import (
	"fmt"
	"testing"
)

func TestChannelEqual(t *testing.T) {
	ch := make(chan int, 0)
	chs := make([]chan int, 0)
	chs = append(chs, ch)

	for i := 0; i < 10; i++ {
		t := make(chan int)
		chs = append(chs, t)

	}

	for _, ints := range chs {
		if ints == ch {
			fmt.Println("equal")
		}
	}
}

func TestChannelEqual2(t *testing.T) {
	chMap := make(map[chan int]string)

	ch := make(chan int)

	chMap[ch] = "1"

	fmt.Println(chMap[ch])
}
