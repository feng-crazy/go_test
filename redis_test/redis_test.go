package redis_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.122.48.17:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	value := []byte{1, 2, 3}
	err := rdb.Set(ctx, "key", value, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Bytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}


func TestRedisPush(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.122.48.17:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})



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
		"f1" : 1,
		"f2" : "2",
	}

	value1 := map[string]interface{}{
		"f1" : 3,
		"f2" : "4",
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

