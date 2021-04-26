package common_test

import (
	"fmt"
	"regexp"
	"runtime"
	"testing"
)

func TestOsInfo(t *testing.T) {
	fmt.Println(runtime.GOOS)
}

func TestStringReplace(t *testing.T) {
	s := `aaaaa
		bbbb
		cc substring ddd
		eeee
		ffff`
	re := regexp.MustCompile("(?m)[\r\n]+^.*b.*$")
	res := re.ReplaceAllString(s, "")
	fmt.Println(res)
}

func TestStringCom(t *testing.T) {
	s := "\""
	fmt.Println(s[0])
	if s[0] == '"' {

	}
}
