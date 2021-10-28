package map_test

import (
	"fmt"
	"testing"

	"github.com/cornelk/hashmap"
)


func TestHashMap(t *testing.T) {
	m := &hashmap.HashMap{}
	m.Set("amount", 123)
	amount, ok := m.Get("amount")
	fmt.Println(amount, ok)
}
