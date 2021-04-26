package cron

import (
	"fmt"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
)

func TestCronBase(t *testing.T) {
	c := cron.New()

	c.AddFunc("@every 1s", func() {
		fmt.Println("tick every 1 second")
	})

	c.Start()
	time.Sleep(time.Second * 5)
}
