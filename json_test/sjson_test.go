package json_test

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/sjson"
)

var testData = ``

func BenchmarkSjson(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = sjson.Set(testData, "name.first", "test")
		_, _ = sjson.Set(testData, "name.last", "Anderson")
	}
}

var person1 = Person{
	Name:     Name{
		First: "test",
		Last:  "Anderson",
	},
}
func BenchmarkJsoniter(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = jsoniter.Marshal(person1)
	}
}