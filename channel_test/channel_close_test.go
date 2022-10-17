package channel

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestClosetChannel(t *testing.T) {
	rand.Seed(time.Now().Unix())

	ch := make(chan int, 10)

	go func() {
		for {
			v := rand.Int()
			// fmt.Println("set v: ", v)
			t, ok := <-ch
			fmt.Println(t, ok)
			ch <- v
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			count := 0
			for i := range ch {
				count++
				fmt.Println("get v:", i)
				if count == 2 {
					close(ch)
					break
				}
			}

			fmt.Println("----------------")
		}
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}

func TestClosetChannel2(t *testing.T) {
	rand.Seed(time.Now().Unix())

	ch := make(chan int, 10)

	go func() {
		count := 0
		for {
			v := rand.Int()
			fmt.Println("set v: ", v)
			ch <- v
			time.Sleep(500 * time.Millisecond)
			count++
			if count > 10 {
				close(ch)
				break
			}
		}
		fmt.Println("--------set end--------")
	}()

	go func() {
		count := 0
		for i := range ch {
			count++
			fmt.Println("get v:", i)
		}

		fmt.Println("--------get end--------")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}
