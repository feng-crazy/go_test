package common

import (
	"sync"
	"testing"

	"github.com/cornelk/hashmap"
)

var data = make(map[string]string)
var sm sync.Map
var m = &hashmap.HashMap{}

func init() {
	data["test"] = "test1"
	sm.Store("test", "test1")
	m.Set("test", "test1")
}

func BenchmarkWithRWMutex(b *testing.B) {
	b.ReportAllocs()
	var mu sync.RWMutex
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = data["test"]
		mu.RUnlock()
	}
}

func BenchmarkNOWithRWMutex(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = data["test"]
	}
}

func BenchmarkWithSyncMap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = sm.Load("test")
	}
}

func BenchmarkWithHashMap(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = m.Get("amount")
	}
}
