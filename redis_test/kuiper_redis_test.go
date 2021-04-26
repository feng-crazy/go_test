package redis_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestRedisPop(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.122.48.17:6379",
		Password: "abc123", // no password set
		DB:       3,        // use default DB
	})

	for {
		val, err := rdb.Get(ctx, "hdf_test").Bytes()
		if err != nil {

		} else {
			fmt.Println("byte: ", val)
			fmt.Println("string", string(val))
		}

		time.Sleep(1 * time.Second)
	}

}

func TestFlag(t *testing.T) {
	m := make(map[string]bool)

	f, _ := m["haha"]

	fmt.Println(f)
}
