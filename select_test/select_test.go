package select_test

import (
	"fmt"
	"testing"
	"time"
)

func TestSelect0(t *testing.T) {
	c := make(chan int)
	c1 := make(chan int)

	c <- 0
	c1 <- 1

	select {
	case i := <- c:
		fmt.Println("send :", i)

	case j := <- c1:
		fmt.Println("send :", j)
	default:
		return
	}

}


func TestSelect(t *testing.T) {
	c := make(chan int)

	go func() {
		for {
			select {
			case data := <-c:
				fmt.Println("one done: ", data)
			}

		}
	}()

	fmt.Println("start")
	i := 0
	for {
		select {
		case c <- i:
			fmt.Println("send :", i)
		}

		i++
		time.Sleep(time.Second)
	}
}
