package map_test

import (
	"fmt"
	"testing"
)

func TestMapNil(t *testing.T) {
	m := make(map[string]interface{})

	v, ok := m["test"]
	fmt.Println(v, ok)
}
