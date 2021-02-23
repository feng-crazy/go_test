package worker_pool_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/xxjwxc/gowp/workpool"
)

func dofunc(i int)(int, error){
	return i, errors.New(strconv.Itoa(i))
}

func TestWorker(t *testing.T)  {
	wp := workpool.New(5) // 设置最大线程数
	fmt.Println(wp.IsDone())
	// wp.DoWait(func() error {
	//	for j := 0; j < 10; j++ {
	//		fmt.Println(fmt.Sprintf("%v->\t%v", 1111111, j))
	//	}
	//
	//	return nil
	//	// time.Sleep(1 * time.Second)
	//	// return errors.New("my test err")
	// })

	errList := make([]error, 0)
	for i := 0; i < 10; i++ { // 开启10个请求
		ii := i
		wp.Do(func() error {
			n, err := dofunc(ii)
			fmt.Printf("%d\n", n)
			errList = append(errList, err)
			return nil
		})

	}

	_ = wp.Wait()
	fmt.Println(wp.IsDone())
	fmt.Println(wp.IsClosed())
	fmt.Println("down")
	fmt.Println(errList)
}
