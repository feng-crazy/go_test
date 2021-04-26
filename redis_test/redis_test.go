package redis_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.122.48.17:6379",
		Password: "abc123", // no password set
		DB:       8,        // use default DB
	})

	value := []byte{1, 4, 3}
	err := rdb.Set(ctx, "key11", value, -1*time.Second).Err()
	if err != nil {
		panic(err)
	}

	// rdb.Expire(ctx, "key", 6* time.Second)
	val, err := rdb.Get(ctx, "key11").Bytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}

func TestRedisPush(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.122.48.17:6379",
		Password: "abc123", // no password set
		DB:       0,        // use default DB
	})

	rdb.Expire(ctx, "keypush", 2*time.Second)

	value1 := []byte{1, 2, 3}
	value2 := []byte{94, 97, 99}
	err := rdb.LPush(ctx, "keypush", value1, value2).Err()
	if err != nil {
		panic(err)
	}

	value3 := []byte{69, 68, 67}

	err = rdb.LPush(ctx, "keypush", value3).Err()
	if err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)

	val, err := rdb.RPop(ctx, "keypush").Bytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("keypush", val)

	val, err = rdb.RPop(ctx, "keypush").Bytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("keypush", val)

	val, err = rdb.RPop(ctx, "keypush").Bytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("keypush", val)
}

func TestRedisRPop(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.122.48.17:6379",
		Password: "abc123", // no password set
		DB:       0,        // use default DB
	})

	val, err := rdb.RPop(ctx, "keypush").Bytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("keypush", val)

	val, err = rdb.RPop(ctx, "keypush").Bytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("keypush", val)

	val, err = rdb.RPop(ctx, "keypush").Bytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("keypush", val)
}

func TestRedisHash(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.122.48.17:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// value := struct {
	// 	f1 string
	// 	f2 string
	// }{"1", "2"}

	value := map[string]interface{}{
		"f1": 1,
		"f2": "2",
	}

	value1 := map[string]interface{}{
		"f1": 3,
		"f2": "4",
	}

	err := rdb.HMSet(ctx, "keyHash", value, value1).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.HGetAll(ctx, "keyHash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("keyHash", val)

}
