package go_cache_test

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

var DevChannel map[string]chan int

func TestMap(t *testing.T) {
	DevChannel = make(map[string]chan int)
	rand.Seed(time.Now().Unix())
	lc := cache.New(time.Minute*5, time.Minute*2)
	log.Printf("start at:%v", time.Now())
	log.Println("set run over")

	for i := 0; i < goroutineNums; i++ {
		go func(idx int) {
			for {
				newKey := fmt.Sprintf("%s:%d", "key", idx)
				c := make(chan int, 1)
				DevChannel[newKey] = c
				c <- idx
			}

		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(lc.ItemCount())

	for i := 0; i < goroutineNums; i++ {
		go func(idx int) {
			for {
				key := fmt.Sprintf("%s:%d", "key", idx)
				c := DevChannel[key]
				select {
				case value := <-c:
					fmt.Printf("key :%s, value:%d", key, value)
				default:
					time.Sleep(1 * time.Second)

				}
			}
		}(i)
	}

	select {}
}

func TestMapDel(t *testing.T) {
	m1 := map[string]string{}
	m1["1"] = "1"
	delete(m1, "2")
	fmt.Println(m1)
}
