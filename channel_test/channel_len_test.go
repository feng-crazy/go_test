package channel

import (
	"fmt"
	"testing"
)

func TestChannelLen(t *testing.T) {
	ch := make(chan int, 10)
	fmt.Println(len(ch))

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4

	fmt.Println(<-ch)
	fmt.Println(len(ch))
}
