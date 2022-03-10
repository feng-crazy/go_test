package compress_test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {
	buf := bytes.Buffer{}

	buf.Write([]byte("test..."))
	fmt.Println(buf.Len())
	fmt.Println(string(buf.Bytes()))

	fmt.Println(buf.String())
	result := make([]byte, 6)

	fmt.Println(buf.Read(result))
	fmt.Println(string(result))
}

func TestBuffer1(t *testing.T) {
	data := []byte("test...")
	buf := bytes.NewBuffer(data)
	fmt.Println(buf.Len())
	fmt.Println(string(buf.Bytes()))
}
