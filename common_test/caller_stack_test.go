package common

import (
	"fmt"
	"runtime"
	"testing"
)

func foo() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(1, pc)
	frames := runtime.CallersFrames(pc[:n])

	var frame runtime.Frame
	more := n > 0
	output := ""
	for more {
		frame, more = frames.Next()
		output = output + fmt.Sprintf("%s:%d \n %s\n", frame.File, frame.Line, frame.Function)
	}
	fmt.Println(output)
}

func foo1() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(1, pc)
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		fmt.Printf("%s %d %s\n", file, line, f.Name())
	}
}

func foo2() {
	pc, file, line, ok := runtime.Caller(0)
	if ok {
		fmt.Println(runtime.FuncForPC(pc).Name(), file, line)
	}
}

func TestCaller(t *testing.T) {
	foo()
	fmt.Println()
	foo1()
	fmt.Println()
	foo2()
}
