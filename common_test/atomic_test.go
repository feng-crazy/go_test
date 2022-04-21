package common

import (
	"fmt"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var tmp atomic.Value
	tmp.Store(false)
	if tmp.Load() == true {
		fmt.Println("true ", tmp.Load())
	}
	if tmp.Load() == false {
		fmt.Println("false ", tmp.Load())
	}
	i := tmp.Load().(bool)
	if i {
		fmt.Println(i)
	}

}

func TestAtomicInt(t *testing.T) {
	var tmp int32
	atomic.StoreInt32(&tmp, 1)
	fmt.Println(atomic.LoadInt32(&tmp))
	atomic.AddInt32(&tmp, 1)
	fmt.Println(atomic.LoadInt32(&tmp))
}
