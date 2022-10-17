package channel_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestChannel1(t *testing.T) {
	rand.Seed(time.Now().Unix())

	ch := make(chan int)

	chs := append([]chan int{}, ch)

	for i := 0; i < 10; i++ {
		t := make(chan int)
		chs = append(chs, t)

	}

	go func() {
		for {
			v := rand.Int()
			fmt.Println("set v: ", v)
			ch <- v
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			v := rand.Int()
			fmt.Println("set v: ", v)
			ch <- v
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		count := 0
		for i := range ch {
			fmt.Println("get v:", i)
			count++
			if count == 2 {
				// close(ch)
				fmt.Println("exit")
				break
			}
		}
	}()

	for _, ints := range chs {
		if ints == ch {
			fmt.Println("equal")
		}
	}

	for {
		time.Sleep(1 * time.Second)
	}
}

func TestChannel(t *testing.T) {
	rand.Seed(time.Now().Unix())

	ch := make(chan int)

	go func() {
		for {
			v := rand.Int()
			// fmt.Println("set v: ", v)
			ch <- v
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			for i := range ch {
				fmt.Println("get v:", i)
			}

			fmt.Println("----------------")
		}
	}()

	select {}
}
