package nux_test

import (
	"fmt"
	"testing"

	"github.com/feng-crazy/go-utils/nux"
)

func TestNux(t *testing.T) {
	fmt.Println(nux.NtpTime("127.0.0.1"))
}
