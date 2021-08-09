package common

import (
	"fmt"
	"testing"
)

var ChannelDevMap = make(map[string]map[string]string)

func TestIf(t *testing.T) {
	ChannelDevMap["1"] = make(map[string]string)
	chlName := "1"
	name := "a"
	if devMap, ok := ChannelDevMap[chlName]; ok {
		devMap[name] = "test_if"
		ChannelDevMap[chlName] = devMap
	} else {
		devMap = make(map[string]string)
		devMap[name] = "test_else"
		ChannelDevMap[chlName] = devMap
	}

	fmt.Println(ChannelDevMap)
}
