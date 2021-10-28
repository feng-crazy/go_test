package json_test

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/gjson"
)

var jsonData = []byte(`{
  "age":37,
  "children": ["Sara","Alex","Jack"],
  "fav.movie": "Deer Hunter",
  "friends": [
    {"age": 44, "first": "Dale", "last": "Murphy"},
    {"age": 68, "first": "Roger", "last": "Craig"},
    {"age": 47, "first": "Jane", "last": "Murphy"}
  ],
  "name": {"first": "Tom", "last": "Anderson"}
}`)


type Person struct {
	Age int `json:"age"`
	Children []string `json:"children"`
	FavMovie string `json:"fav.movie"`
	Friends []Friends `json:"friends"`
	Name Name `json:"name"`
}
type Friends struct {
	Age int `json:"age"`
	First string `json:"first"`
	Last string `json:"last"`
}
type Name struct {
	First string `json:"first"`
	Last string `json:"last"`
}

func BenchmarkDecodeJsoniterPersonGetName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = jsoniter.Get(jsonData, "name", "last").ToString()
		// fmt.Println(value)
	}
}

func BenchmarkDecodeGjsonStructPersonGetName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ =gjson.GetBytes(jsonData, "name.last").String()
		// fmt.Println(value)
	}
}