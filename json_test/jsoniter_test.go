package json_test

import (
	"fmt"
	"testing"

	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

func TestJsoniterGet(t *testing.T) {
	should := require.New(t)
	any := jsoniter.Get([]byte(`{"a":{"stream":{"c":"d"}}}`))
	fmt.Println(any.ToString())
	should.Equal("d", any.Get("a", "stream", "c").ToString())
}

func TestJsoniterGet1(t *testing.T) {
	any := jsoniter.Get([]byte(`{"a":{"stream":{"c":"d"}}}`), "a", "stream", "c")
	fmt.Println(any.ToString())
}
