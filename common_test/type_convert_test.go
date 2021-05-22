package common

import (
	"fmt"
	"strconv"
	"testing"
)

func Benchmark_NormalConvert(b *testing.B) {
	var a int64 = 100

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(a, 10)
	}
}

func Benchmark_AssertConvert(b *testing.B) {
	var a int64 = 100
	for i := 0; i < b.N; i++ {
		ToStringE(a)
	}
}

func ToStringE(i interface{}) (string, error) {
	switch s := i.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case []byte:
		return string(s), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", i, i)
	}
}
