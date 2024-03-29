package common

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/spf13/cast"
)

func TestByteToString(t *testing.T) {
	var b = []byte{123, 34, 101, 118, 101, 110, 116, 68, 97, 116, 97, 34, 58, 34, 50, 48, 50, 49, 45, 48, 57, 45, 48, 56, 32, 49, 49, 58, 51, 53, 58, 48, 50, 34, 44, 34, 115, 116, 97, 116, 105, 111, 110, 83, 116, 97, 116, 117, 115, 34, 58, 34, 50, 34, 44, 34, 115, 116, 97, 116, 105, 111, 110, 73, 110, 102, 111, 34, 58, 34, 49, 52, 45, 50, 228, 184, 128, 229, 177, 130, 232, 165, 191, 230, 142, 167, 229, 136, 182, 229, 153, 168, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 48, 34, 44, 34, 100, 101, 118, 105, 99, 101, 83, 121, 115, 78, 111, 34, 58, 34, 51, 57, 34, 125}
	fmt.Println(string(b))

	b = []byte{123, 34, 101, 118, 101, 110, 116, 68, 97, 116, 97, 34, 58, 34, 50, 48, 50, 49, 45, 48, 57, 45, 48, 56, 32, 49, 51, 58, 51, 51, 58, 51, 48, 34, 44, 34, 100, 111, 111, 114, 83, 116, 97, 116, 117, 115, 34, 58, 51, 44, 34, 100, 111, 111, 114, 78, 97, 109, 101, 34, 58, 34, 49, 50, 35, 228, 184, 128, 229, 177, 130, 228, 184, 156, 229, 144, 136, 231, 148, 168, 229, 137, 141, 229, 174, 164, 34, 44, 34, 100, 111, 111, 114, 73, 100, 34, 58, 54, 53, 44, 34, 117, 114, 103, 101, 110, 116, 84, 121, 112, 101, 34, 58, 50, 44, 34, 101, 118, 101, 110, 116, 84, 121, 112, 101, 34, 58, 50, 125}
	fmt.Println(string(b))
}

func TestByteToString1(t *testing.T) {
	b := []byte{123, 34, 100, 97, 116, 97, 34, 58, 34, 123, 92, 34, 97, 108, 97, 114, 109, 95, 103, 114, 97, 100, 101, 92, 34, 58, 92, 34, 53, 92, 34, 44, 92, 34, 97, 108, 97, 114, 109, 95, 116, 121, 112, 101, 92, 34, 58, 92, 34, 86, 77, 83, 45, 80, 85, 45, 65, 76, 65, 82, 77, 92, 34, 44, 92, 34, 97, 108, 97, 114, 109, 95, 116, 121, 112, 101, 95, 111, 98, 106, 92, 34, 58, 123, 92, 34, 100, 101, 115, 99, 114, 105, 112, 116, 105, 111, 110, 92, 34, 58, 92, 34, 232, 174, 190, 229, 164, 135, 233, 149, 191, 230, 151, 182, 233, 151, 180, 231, 166, 187, 231, 186, 191, 229, 145, 138, 232, 173, 166, 92, 34, 44, 92, 34, 101, 118, 101, 110, 116, 95, 116, 121, 112, 101, 92, 34, 58, 92, 34, 49, 54, 57, 48, 56, 50, 57, 49, 92, 34, 44, 92, 34, 103, 114, 97, 100, 101, 92, 34, 58, 53, 44, 92, 34, 105, 100, 92, 34, 58, 51, 48, 44, 92, 34, 109, 113, 92, 34, 58, 48, 44, 92, 34, 111, 98, 106, 92, 34, 58, 51, 44, 92, 34, 114, 101, 109, 97, 114, 107, 92, 34, 58, 92, 34, 232, 174, 190, 229, 164, 135, 233, 149, 191, 230, 151, 182, 233, 151, 180, 231, 166, 187, 231, 186, 191, 229, 145, 138, 232, 173, 166, 92, 34, 44, 92, 34, 115, 111, 117, 114, 99, 101, 95, 116, 121, 112, 101, 92, 34, 58, 92, 34, 86, 77, 83, 45, 80, 85, 45, 65, 76, 65, 82, 77, 92, 34, 44, 92, 34, 115, 116, 97, 116, 92, 34, 58, 48, 44, 92, 34, 115, 116, 111, 114, 109, 95, 116, 105, 109, 101, 92, 34, 58, 49, 48, 44, 92, 34, 116, 121, 112, 101, 95, 110, 97, 109, 101, 92, 34, 58, 92, 34, 69, 95, 80, 85, 95, 79, 102, 102, 108, 105, 110, 101, 76, 111, 110, 103, 116, 105, 109, 101, 65, 108, 97, 114, 109, 92, 34, 125, 44, 92, 34, 97, 108, 97, 114, 109, 95, 116, 121, 112, 101, 95, 115, 117, 98, 92, 34, 58, 92, 34, 69, 95, 80, 85, 95, 79, 102, 102, 108, 105, 110, 101, 76, 111, 110, 103, 116, 105, 109, 101, 65, 108, 97, 114, 109, 92, 34, 44, 92, 34, 99, 104, 97, 110, 110, 101, 108, 95, 110, 97, 109, 101, 92, 34, 58, 92, 34, 229, 164, 185, 229, 177, 130, 231, 131, 173, 229, 138, 155, 229, 176, 143, 229, 174, 164, 232, 181, 176, 233, 129, 147, 92, 34, 44, 92, 34, 100, 101, 118, 105, 99, 101, 95, 110, 97, 109, 101, 92, 34, 58, 92, 34, 229, 164, 185, 229, 177, 130, 231, 131, 173, 229, 138, 155, 229, 176, 143, 229, 174, 164, 232, 181, 176, 233, 129, 147, 92, 34, 44, 92, 34, 103, 97, 95, 99, 111, 100, 101, 92, 34, 58, 92, 34, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 49, 51, 50, 49, 48, 48, 48, 50, 51, 55, 92, 34, 44, 92, 34, 105, 100, 92, 34, 58, 92, 34, 56, 56, 53, 49, 57, 56, 53, 51, 56, 57, 50, 56, 54, 56, 57, 49, 53, 50, 92, 34, 44, 92, 34, 111, 114, 103, 95, 105, 110, 100, 101, 120, 92, 34, 58, 92, 34, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 50, 49, 54, 48, 54, 48, 48, 48, 49, 53, 92, 34, 44, 92, 34, 111, 114, 103, 95, 110, 97, 109, 101, 92, 34, 58, 92, 34, 49, 52, 45, 50, 230, 165, 188, 92, 34, 44, 92, 34, 112, 117, 95, 105, 100, 92, 34, 58, 92, 34, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 49, 51, 50, 49, 48, 48, 48, 50, 51, 55, 92, 34, 44, 92, 34, 114, 101, 112, 111, 114, 116, 95, 116, 105, 109, 101, 92, 34, 58, 49, 54, 51, 49, 48, 56, 57, 51, 55, 55, 48, 48, 48, 44, 92, 34, 115, 116, 97, 116, 117, 115, 92, 34, 58, 92, 34, 48, 92, 34, 125, 34, 44, 34, 101, 118, 101, 110, 116, 95, 116, 121, 112, 101, 34, 58, 34, 48, 48, 49, 48, 48, 51, 34, 44, 34, 101, 118, 101, 110, 116, 95, 105, 100, 34, 58, 34, 54, 56, 51, 52, 52, 101, 100, 56, 98, 53, 100, 48, 52, 56, 49, 56, 57, 50, 52, 55, 99, 53, 99, 54, 53, 99, 54, 98, 100, 50, 56, 54, 34, 125}
	fmt.Println(string(b))
}

func TestFloadToBool(t *testing.T) {
	f := 1.111
	var b bool

	b = cast.ToBool(f)
	fmt.Println(b)
}

func TestInterfaceNil(t *testing.T) {
	var p *int = nil
	var v interface{} = p
	fmt.Println(p == v)
	fmt.Println(v == nil)
}

type Link struct {
	value int
	Next  *Link
}

func PrintLink(head *Link) {
	for head.Next != nil {
		fmt.Println(head.value)
		node := head.Next
		head = node
	}
}

func TestTmp1(t *testing.T) {
	l := new(Link)
	head := l
	for i := 0; i < 5; i++ {
		node := &Link{
			value: i + 1,
			Next:  nil,
		}
		head.Next = node
		head = node
	}
	PrintLink(l)
	left, right := 2, 4
	PrintLink(Revert(l, left, right))
}

func Revert(head *Link, left, right int) *Link {
	if head == nil || left >= right {
		return head
	}
	newHead := &Link{value: 0, Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < left-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		tmp := pre.Next
		pre.Next = tmp.Next
		cur.Next = pre.Next
		pre.Next = tmp
	}
	return newHead.Next
}

type sc struct {
	left  int
	right interface{}
	sm    string
}

func calc(str string) int {
	var result int

	data := []byte(str)
	for _, v := range data {
		if v >= '0' && v <= '9' {

		} else if v == '+' || v == '-' || v == '*' || v == '/' {

		} else if v == '(' {
			// 入栈
		} else if v == ')' {
			// 出栈
		}
	}
	return result
}

func TestTmp(t *testing.T) {

}

var done chan struct{}

func sender(ctx context.Context, ch chan int) {
	for {
		// math.Rand
		select {
		case <-ctx.Done():
			return
		default:
			ch <- rand.Int()
			time.Sleep(1 * time.Second)
		}
	}
}

func recv(ctx context.Context, ch chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Println("recv: ", v)
		case <-ctx.Done():
			return
		}
	}
}

func recv1(ctx context.Context, ch chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("recv%v, %d: ", ctx.Value("num"), v)
		case <-ctx.Done():
			return
		}
	}
}

func TestChannel(t *testing.T) {
	num := 3
	ch := make(chan int, num)

	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), "num", num))

	go sender(ctx, ch)
	for i := 0; i < num; i++ {
		go recv(ctx, ch)
	}

	// go recv1(ctx, ch)
	time.Sleep(5 * time.Second)
	cancel()
}

func sliceHandle(rv <-chan string, sd chan<- map[byte]int) {
	m := make(map[byte]int)
	for v := range rv {
		ss := []byte(v)
		for _, sv := range ss {
			m[sv]++
		}
	}
}

var wait = sync.WaitGroup{}

func mapHandle(sd <-chan map[byte]int) {
	m := make(map[byte]int)
	for {
		select {
		case r := <-sd:
			for k, v := range r {
				m[k] = v
			}
			wait.Wait()
		}

	}
	fmt.Println(m)
}

func TestSlice(t *testing.T) {
	workNum := 3

	rv := make(chan string, workNum)
	sd := make(chan map[byte]int, workNum*2)
	s := []string{"abc", "cbd", "ebg"}

	wait.Add(1)
	go mapHandle(sd)

	for i := 0; i < workNum; i++ {
		go sliceHandle(rv, sd)
	}

	for _, v := range s {
		rv <- v
	}

	wait.Done()
}
