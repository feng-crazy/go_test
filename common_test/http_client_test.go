package common_test

import (
	"fmt"
	"testing"

	"gopkg.in/resty.v1"
)

func TestHttpClient(t *testing.T) {
	resp, err := resty.R().SetPathParams(map[string]string{
		"namespace": "xuanyuan",
		"app_name":  "xuanyuan-middle-xuanyuan"}).Get("http://10.122.130.254:17778/api/v1/app/{namespace}/{app_name}/devices")

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	// fmt.Println(resp)
}
