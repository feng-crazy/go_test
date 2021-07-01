package go_cache_test

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
)

var goroutineNums = 10

func TestCache(t *testing.T) {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	rand.Seed(time.Now().Unix())
	lc := cache.New(time.Minute*5, time.Minute*2)
	log.Printf("start at:%v", time.Now())
	log.Println("set run over")

	for i := 0; i < goroutineNums; i++ {
		go func(idx int) {
			newKey := fmt.Sprintf("%s:%d", "key", idx)
			v := rand.Int()
			lc.Set(newKey, v, 2*time.Second)

			fmt.Printf("set vlaue key %s, value %d\n", newKey, v)
		}(i)
	}
	fmt.Println(lc.ItemCount())
	fmt.Println(lc.Items())
	time.Sleep(1 * time.Second)

	// 查看 go-cache 中 key 的数量
	for i := 0; i < goroutineNums; i++ {
		go func(idx int) {
			for {
				key := fmt.Sprintf("%s:%d", "key", idx)
				v, ok := lc.Get(key)
				if ok {
					fmt.Printf("get vlaue key %s, value %+v\n", key, v)
				} else {
					fmt.Println("get value false:")
					break
				}
				time.Sleep(1 * time.Second)
			}
		}(i)
	}

	// // 查看 go-cache 中 key 的数量
	// for i := 0; i < goroutineNums; i++ {
	// 	key := fmt.Sprintf("%s:%d", "key", i)
	// 	fmt.Printf("get vlaue key %s\n",key)
	// 	v, ok := lc.Get(key)
	// 	if ok{
	// 		fmt.Printf("get vlaue key %s, value %+v\n",key, v)
	// 	}else {
	// 		fmt.Println("get value false:")
	// 		break
	// 	}
	// }
}

func TestCache1(t *testing.T) {
	var gatewayStatusCache = cache.New(3*time.Second, 1*time.Minute)
	gatewayStatusCache.Set("test", true, 3*time.Second)
	value, ok := gatewayStatusCache.Get("test")
	fmt.Println(value, ok)
	time.Sleep(4 * time.Second)

	value, ok = gatewayStatusCache.Get("test")
	fmt.Println(value, ok)

	gatewayStatusCache.Set("test", false, 3*time.Second)
	value, ok = gatewayStatusCache.Get("test")
	fmt.Println(value, ok)
	time.Sleep(4 * time.Second)

	value, ok = gatewayStatusCache.Get("test")
	fmt.Println(value, ok)

}
