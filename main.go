package main

import "fmt"

var (
	Version     = "<UNDEFINED>"
	BuildTime   = "<UNDEFINED>"
	VersionType = "<UNDEFINED>"
)

func main() {
	fmt.Println(Version, BuildTime, VersionType)
}
