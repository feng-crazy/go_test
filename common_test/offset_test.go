package common

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOffsetOpration(t *testing.T) {
	fmt.Println(1 << 8)
	fmt.Println(1 >> 8)
}

func TestMacAddrConvert(t *testing.T) {
	macStr := "08:60:6e:fe:47:90"
	macStrArr := strings.Split(macStr, ":")
	fmt.Println(macStrArr)
	macStr1 := strings.Join(macStrArr, "")
	macArr, err := hex.DecodeString(macStr1)
	if err != nil {
		logrus.Error(err)
	}

	fmt.Println(macArr)

	macAddr := fmt.Sprintf("%02x:%02x:%02x:%02x:%x:%02x", macArr[0], macArr[1], macArr[2], macArr[3], macArr[4], macArr[5])
	fmt.Println(macAddr)
}
