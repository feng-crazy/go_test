package common_test

import (
	"fmt"
	"regexp"
	"testing"
)

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
