package common

import (
	"fmt"
	"testing"
)

func TestCn(t *testing.T) {
	str := "中文"

	for _, i2 := range str {
		fmt.Println(i2)
	}

	fmt.Println("\n*********************")
	b := []byte(str)

	for _, b2 := range b {
		fmt.Println(b2)
	}

	fmt.Println("\n*********************")
	for _, b2 := range b {
		fmt.Printf("%d\n", b2)
	}

	fmt.Println("\n*********************")
	b = []byte{228,
		184,
		173,
		230,
		150,
		135}
	fmt.Printf("%s\n", string(b))
}
