package time

import (
	"fmt"
	"testing"
	"time"

	"github.com/rfyiamcool/go-timewheel"
)

func checkTimeCost(t *testing.T, start, end time.Time, before int, after int) bool {
	due := end.Sub(start)
	if due > time.Duration(after)*time.Millisecond {
		t.Error("delay run")
		return false
	}

	if due < time.Duration(before)*time.Millisecond {
		t.Error("run ahead")
		return false
	}

	return true
}

func TestTimeWheel(t *testing.T) {
	tw, _ := timewheel.NewTimeWheel(100*time.Millisecond, 60)
	tw.Start()
	defer tw.Stop()

	queue := make(chan time.Time, 2)
	task := tw.AddCron(time.Second*1, func() {
		queue <- time.Now()
	})

	time.AfterFunc(5*time.Second, func() {
		tw.Remove(task)
	})

	exitTimer := time.NewTimer(10 * time.Second)
	lastTs := time.Now()
	c := 0
	for {
		select {
		case <-exitTimer.C:
			if c > 6 {
				t.Error("cron stop failed")
			}
			return

		case now := <-queue:
			c++
			checkTimeCost(t, lastTs, now, 900, 1200)
			fmt.Println("time since: ", now.Sub(lastTs))
			lastTs = now
		}
	}
}

func TestTimeWheel1(t *testing.T) {
	tw, err := timewheel.NewTimeWheel(500*time.Millisecond, 100)
	if err != nil {
		panic(err)
	}

	tw.Start()
	defer tw.Stop()
	start := time.Now()
	// tw.Add(time.Duration(498*time.Millisecond), func() {
	// 	fmt.Println("Add timer cost: ", time.Since(start))
	// })
	//
	// start = time.Now()

	task := tw.AddCron(1*time.Second, func() {
		fmt.Println("AddCron timer cost: ", time.Since(start))
	})

	fmt.Println(task)
	select {}
}
