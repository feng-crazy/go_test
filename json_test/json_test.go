package json_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type peerInfo struct {
	HTTPPort int
	TCPPort  int
	Other    string
	Version  string
}

func TestJsonTag(t *testing.T) {
	pi := peerInfo{80, 3306, "0.0.1", ""}
	js, err := json.Marshal(pi)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(js))
}

func TestJsonTag1(t *testing.T) {
	var v peerInfo
	data := []byte(`{"http_port":80,"TCPPort":3306}`)
	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", v)
}
