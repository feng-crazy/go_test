package time_test

import (
	"fmt"
	"testing"
	"time"
)

func TestTiker(t *testing.T) {
	MyTimer()
	for {
		time.Sleep(time.Second)
		fmt.Println("main .....")
	}
}

func MyTimer() {
	tiker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-tiker.C:
				fmt.Println("Time now:", time.Now().Format("2006-01-02 15:04:05"))
				tiker.Reset(1 * time.Second)
			}
		}
	}()
}
