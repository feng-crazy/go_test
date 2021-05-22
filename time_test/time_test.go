package time_test

import (
	"fmt"
	"testing"
	"time"
)

type Input struct {
	Timestamp int64 `json:"timestamp"`
}

func InputFilter(timestamp *int64) bool {
	ts := *timestamp

	y := ts % (1000 * 60 * 10)

	// 约定必须一分钟采集一次
	if y <= (1000 * 60) {
		// 使用十分钟整数采集
		*timestamp = ts - y
		return false
	}
	return true
}

func TestTimeRemainder(t *testing.T) {
	for {
		var input Input
		input.Timestamp = time.Now().UnixNano() / 1e6

		if InputFilter(&input.Timestamp) {
			fmt.Println("Filter.........")

			tm := time.Unix(input.Timestamp/1e3, 0)
			fmt.Println(tm.Format("2006-01-02 15:04:05"))
		} else {
			tm := time.Unix(input.Timestamp/1e3, 0)
			fmt.Println(tm.Format("2006-01-02 15:04:05"))
			return
		}

		fmt.Println("----------")

		time.Sleep(1 * time.Minute)
	}

}

func TestTimeTicker(t *testing.T) {
	tiker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-tiker.C:
				fmt.Println("tiker.....", time.Now())
				tiker.Reset(time.Second * 2)
			}
		}
	}()
	select {}
}
