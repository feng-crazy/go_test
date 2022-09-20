package main

import (
	"fmt"
	"time"
)

var (
	Version     = "<UNDEFINED>"
	BuildTime   = "<UNDEFINED>"
	VersionType = "<UNDEFINED>"
)

func main() {
	fmt.Println(Version, BuildTime, VersionType)
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println(i)
		}()
		time.Sleep(time.Second)
	}
}
