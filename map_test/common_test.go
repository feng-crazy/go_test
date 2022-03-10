package map_test

import (
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestCommon(t *testing.T) {
	var itemMap map[string]interface{}

	j := `{"test":1, "metadata":null}`

	_ = jsoniter.UnmarshalFromString(j, &itemMap)

	tmp, _ := itemMap["metadata"].(map[string]interface{})
	fmt.Println(tmp)
}
