package common

import (
	"fmt"
	"testing"
)

func TestForContinue(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	arr2 := []int{2, 4}

	arr3 := []int{}
for1:
	for _, i := range arr {

		for _, i2 := range arr2 {
			fmt.Printf("i:%d, i2:%d\n", i, i2)
			if i == i2 {
				arr3 = append(arr3, i)
				fmt.Println("continue")
				continue for1
			}

		}
	}

	fmt.Println(arr3)
}
