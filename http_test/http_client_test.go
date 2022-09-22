package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"

	"github.com/wenzhenxi/gorsa"
)

func TestHttpClient(t *testing.T) {
	client := resty.New()
	resp, err := client.R().SetPathParams(map[string]string{
		"namespace": "xuanyuan",
		"app_name":  "xuanyuan-middle-xuanyuan"}).Get("http://10.122.130.254:17778/api/v1/app/{namespace}/{app_name}/devices")

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	// fmt.Println(resp)

	// resp, err := resty.R().
	// 	SetHeader("Content-Type", "application/json").
	// 	SetBody(body).
	// 	SetPathParams(map[string]string{
	// 		"namespace": viper.GetString("namespace"),
	// 		"app_name":  viper.GetString("app_name"),
	// 		"name":      deviceName,
	// 	}).
	// 	Post(url)
}

func TestHttpClient1(t *testing.T) {
	data := map[string]string{
		"username":     "admin",
		"password":     "0192023a7bbd73250516f069df18b500",
		"rememberName": "true",
		"execution":    "ca308d5a-807d-443e-ac1e-e99b347ea88a_ZXlKaGJHY2lPaUpJVXpVeE1pSjkuUDVEMHNVR2IwcndKMWpaa2orUGEzZFdUR29PaFMrYXZzeWMvZ0NHcWI1ZTEzSStzQlVmcUc4Yjh5aXZQRnB5QVRnc0pNSTJRN3FmckVKdllDRlFpS3VYSHRKTHN3dkxBbjJpaGh5Zy9BVE9BTkFrQzFkb3lLK2VzUDBBaVhrZnVnaEZSRGtCUWprWVo2emI5NTVETEZzOXkrYW5Jc0xkT0dIWVNlUjlveUxucWNGWVdTNjNDZUx5MGRad3ZMc25nQmYvK2I5TGFVN2FRcGkvYVU5WDN2d2l2SXJpNkRQSEpxdzJueU5aVS9xQ0R6T1hsSnBrUFFGcFVwL1d3dzVHTlFWeDV2Q1Y3WlBETkU3VmRhQmwzR1pvV3NIWjk0ajI1SUlnNkUrNENHWnh2cmVvcnZ6M0w5QkdHTXFhc090NU5IdHdPYWFSSkVQUENncWdCa04rTFY3bit1b3BxZHR5ZU1iRWZsZGJYS1B5Ni9vaUFIZFIxbUU2UlUwOHczUU1hQy9SdDJZMXBCaGdVdEM5V2h3S2hHOEpiUkkrbnBBYkZBNXNKbWYvT09vQy9zQjh4SW9aa1BraVpFZzgyaFN0VDBOeHIvUnRDWlF1aENiREx4WDlNYnpacWVucVFuK0x4WWdZVFY4SUlrQmJLOFN0VWhpZjNXMnltM1lNeWpyN3BqYTJmQnBKenMzTndMMyszVS8xSVNCa0p0Wm5nellSUjFZRWtzNHlXUTZ6NEJ3eG5hSmhRblpqanI3VzVQMHZtaklPUWtFcTVFUS8xWWpDVVJxRndETDVubjJudUk5U1FaUnJUYmRtQlpoY2Q4SXpMb0htMS9yN1EvYnlsc0dhNWR2ODRLRmhqWGgreGk5WE1icWplaEFNSENzandUR1Y4M1B2RDhadllsd3NVRjlhRGpIeVJZSjFqTVN5QXFuWm5KYTNISU1WRzViZUFHTVljWjNOSnJZaldmNzU2bTZmT0xMYmFEc3JjQWF6RnFCN28yZUtnTmVoditYZnFxamtBYk9IeFRjMjVqZVUvNG5qY01JVklYQmFibDd2RmZleTZMUHlGVTJLMk8ycENTeDhUaFd6WExmUnAzZ2hPVjJqVitienBlQlFnN0dPeGRVYy9OZmhSWklKWEJEWUZjSmZvVk5wUDFOVUNQbEludUN4MWJ5Rmlpa0NoS3UrOGEwT1dPTTdmNmErL2dTbnZEbkVVaGhlN2ZQbDViVEZLLzA3UE1Dd09YalZ0dUw4UWdOZGhYTVVlUy9IU1JqQW95UWc5cFBDR2FYOTh3WEx6bVVvNUJ3Smk2T3pHUlVkQVNaRUdNVjRDeE42MFljbnhGOUJKMkkybmxDL2hSNm9WMVdPV1B3cjViR3RJYlc3SUMzbXFTTENQL0NyMnd2bHV3RzRKdDU0K3F2NjJNeDZuRkU5VUorcXBWck5EVGs0SythbmVxekFrc3NZd2ZHR1h5NEpwM0x2VDdOMVBzSFdDUFpQWG1CYTY0OXkzamJ2YmMxdkxxMW9RQ2pDNVpRWDZRaE9uOUJnc2xNTzl0WExoVHNQS0lleS9yVDhZamZSOFAwTlRYQzMzMktSZGJQN0tmTnI1a2ljbDYyNGVRZ1pyV3lMKzU4UEVHLzZjS25Yb2dRUG5BMHdlYWFFaG9MRDdacnF3Slc1ZHE3eGoyYnFzcFhIWkp2NDF3d3JIaVgwc09RSUJsTUp1bFNwVHpsZ05mQUM2RnlWYWVZeFBKZkJTMENScHozNTV6d1pkQlNZa3J5TEVISGV6YmdLUXJmbkNDcVJQWXNZcDFxRFRHQTBZR1F4Qk16d29OaGI1WVVNS0FQdHFoTUQ2U0NGL2FrNE9yKzEvNjU5MTRTRnAzZTV6ZXBhanV2MmxKR2I0ek5YOFRZWEErbGtlcEMzK2dScGswbVhWTGFvWmpZQnVMcHZUTm15cWkxeC9FS0xKWXRVeW5VeU5SczFvZ0RQanZiVkZmSTJ6U0x2WXUrMm9yWWhXbm5ySjRxS1FsTTRtMk1uVUY5RHMvRjlXeW40Z3Z4cEdwQXJLS2c2TmtSNktjdno1YlJNU3gvYThYM2RlNEpkRDMwUTlKNHZlRXBucHJxcjZXVGJpNVdzNC9WNTMxS1RWRkR5ZGRxL0NjRzJGTXZuTnBWY3dEMWxZVU5PMHl5Rmo4d00xb3dROHN6VE5LV3pWYlZTQ1dxanNGQUVmQXIvRzJ1OWl1NUIycFFCWmxiSi9pQ1pFVEROTkhXYm9peGNvN3RMaDRZVDZ1UDRsZW9mdSsvZXlaWWNyd0hkY0tMdkxyaDlUeXNFRk1wNTRLMTdnZGJScXVhd1ZUczVuM2p1SXZjYWoyR3NSNHJiUFNjSlFsK09LbU1UK3VVTDhnZHFyTGN2ZytESlRYeG1oOGloaGZ0RnlWTFlqZnkzKzhUQzFFTzdGL1haUEY3ZFgzalpEaGRxc0MxNnkxUkdIYUhWY0VKZ1hjKzJONk9KbkJjUE8yekNyYjNiR0JCaWl0MzN6bW8wYzhZZjUrR0hGa2l3blFIS2JWRU5vZGJZeGttaEhWT2JEK2VBM2xTaEQ0eVltREVsRnRWRk5GWDZFOFd6WmVyK0NKODhFaVNMR3RhNDgrLzJOTkJpOGh5UXBGRDB1YlNWTHBXUVV0YWJ6RFh2ckRuVnRHMkhOQ3BNVWtvN0VtK0h6UXA0R3c4cWwzVlN4b1Y1cXhFMTUxQ3NLWERMei81TGdtTnFLRXBxaENpOXhidGp5VXNqTjFLNFdpUHd2aFBqUnlRZ2xybHZXTTFUWExkYnVValh5ekxxZFVxQ0h4WW1oZ1BxVm1uY3VUODMyVk5pMm9SeDU5ek9QRGFzbU1vZllzNnBwNHY2NXYzMVVxM1BPdlYzMEZWU0tKTEV0bkdVVGI5YWh3UGZiLzErMVplb3BWbFRUQjR3Vnc0cnJaUFNJZktqS1kvRG1sL0pDR294enFKc0lud2F3bFBsblNOdjJ1UlZ3Q3lUQkw5UmdHYStkV20vWlJnYlAyU1F3NUNKbXliQVlCOGVFZHlZUyswbEFuaDhIY0tlSnhWT00zYTVOM011Yng2MzRTQjVocDlMMVpCMEEwSXd2cS8yaXVXd0ZvREo0enBOeG9hOGRlcTJGaXdIUUZBNFV4MGd5bGJXK1VkenRoZUpKbXBKZjhiRUs1TU9tY3dGdE5NUkY1ZHY5UzRqcDVhT0YwSS9QbHF4YXpOeWo2bm5xRzlvRmFqc1lsTnFta2QvbWhWalE0dTlORm10WnBCRXBieEZRWk8xNGZxR2FJbWMxWU9FY0hSYU94N1phcHBPU015NlkrcG5wSmw4UGwvdnhhOHZiZTZERDBwdlN6R0JCU2FtYjR2V1NSbDFONmd5Q0MxNFNxZFcreWIwRjRvbGY1U3dRcXBiR0dEQTgxVlU1ejVUZVFmakR6UXVTRUc2UnhpMjJVakNaN08zR1FiMUVFWTNBeTVzeXVqZjM5Q0RaMkNyRExnVTM2UWJwU2ozbGdZN2lPbmhpeW9IMEtjMTBRSGVoOTZScDdMei9CUWFYQXZnYi9PSjEvbDBMZnhrWSt6ZlYyVXFKKzdJRCtuZWxXSiticjNrM1cwdGZodDM5S2tpSTlZZWNCNUNPdWY0NzdpaTljSERNemZLTTUrTm8xS3BscEluZkVzUDZIdEFDVDlVdEtIWXg0WTBwakw4ODRhRnFEN1Rhc21CNHdINDVFdldqUHcyWWJ4a0NHSGE0T3hqQVlpN3BsTXZhTTVZRG9qNUttY29FRE1YSU1lNDlPeXRUVTd6d2VxSUt1UXlCUGIzYzFtN1Q0RTB6aUhCdEdYSG4zZW5OTmh0SmszMU1DNHltOU9ZUU5kRkZEU1pMKzJLZDYwREw5cUkvY1Z2OVozUDdOUTB3NEFpRlNldVJmSUtBUUxsYnUzVDliSW1OYjNWM1M0V3ZlVDhZOWs4TUJWNktpK2VjMnA5ZmNCd0ZQVnFlcFNnaWhMSEM0eUZuOS9pWTN4N1d0WUxPZlpLdnpKM2RKZWRjUCsreHFrNmIwVFVqcWM5QU1BYlhwRklRUU5IN2t1M2NhQU5RVU1TKzZIWUtDVjlyYUc4eHBpMndtTEl6NUkwVFRLM21LZllhOXFrYWxpQXBqRzgyRStVOFlYMzMxcEZia241VjVwQlBBSU8wVEt4d1I3T2kvZVF3STdHN3Body85QTBaZUdYOVFOZVpwelpvSXhLcE9lQ0tZQlF1RXIzZE93Nk0yRnhqRTZJdVczUWVuTWYwY2E2Mkd6MXh4RTNjYUdjYU5GTHBaYk9XRFRxTzZBdElVVGpWU2RKTDhEY3dTVXRIRk8rd0d1T3FDdDlaeW96RlVpUzVCeEd2RFFxemFRY2VFbGRoM2EwZFc1RXl1eHdGVUhzdUJxZ0UrbWNvQ21TV3ByTTNsbHlXaGJpVUtzMENySU5jMHo5Wit0dWcvU3dSQzhxOXg1QU5HMnRKM1pZajFYOUpLbkxlYjZrcUJiWlUyamV6QjZzZkU0YTJlbWFmdXR1TGk5THRHak5CK0VmWFQvZ3laanUyTCtjd1hLU0lVWEtvV1Z6YWJrczEzdXc4RDdNVlA3UDNURi9oMFd5MWdrVEF5Q3U0Szg1dHNCd0cwbkgrRW81cXVoRzBNT0VtRkg0TzZvV3dDL21MbnBmQ0NlendkakNDcCt3bktpUngxTzVsT3dLbFBORGFSenI0c2RLQ2VNNEZGRWZOSDZqTFIzczVLaElqZnRpUHFzb1c5UWloWWI3QTVPVDh1TUFZd3hUZ0pIZlptUU1YTFM5M3cweFRuYkxWaWc3N2RtbmJUT0RiWG9CblJnTUFJY1kvMTF4Z0xnaUhTN0Y4M291RUJycmdhZGRteTBWSzdzREtTK09IYXpXUUVaRk1sKzdNZzJ6S00waXRhR2pxZm1ITXJ2c2dKSm5qb3VDaUFEMXA1YmRKMDV3OFZPWFNad0VDZUhTS0xrQk1NNGhNWjkwd1Fva2psQlU9LkRHXy0yRWoyNVExdGEySXZtZ3NmOXpGdUhrMjZscG9CckdYRzU2VVozTkdPeTB3SmR0YUgtdG1aTVhETGpqSXhFdWFqc1EzcjZ5dlFfbmZQZVlvb0FB",
		"_eventId":     "submit",
		"submit":       "登录",
	}
	client := resty.New()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))
	resp, err := client.R().SetFormData(data).
		Post("http://10.170.208.212:40087/sso/login?service=http%3A%2F%2F10.170.208.212%3A9000%2Fcas%2Fvalidate&sn=undefined")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Header())
	fmt.Println(client.Cookies)
	fmt.Println(resp.Cookies())
	fmt.Println(resp.RawResponse.Cookies())
	fmt.Println("-------")

	client1 := resty.New()
	client1.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))
	client1.Header = resp.RawResponse.Header
	client1.Cookies = resp.RawResponse.Cookies()
	// client1.Cookies = client.Cookies

	resp, err = client1.R().Get("http://10.170.208.212:9000/api/uss/v1/ams/alarm/type/page?pageNo=1&pageSize=99999")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	fmt.Println("-------")
	fmt.Println(resp.Header())
	fmt.Println(client.Cookies)
	fmt.Println(resp.Cookies())
	fmt.Println(resp.RawResponse.Cookies())
	fmt.Println("-------")
}

func TestHttpClient2(t *testing.T) {
	client := http.Client{}
	str := ""
	reqBody := bytes.NewBuffer([]byte(str))

	req, err := http.NewRequest("POST",
		"http://10.170.208.212:40087/sso/login?service=http%3A%2F%2F10.170.208.212%3A9000%2Fcas%2Fvalidate&sn=undefined",
		reqBody)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

func TestHttpClient3(t *testing.T) {
	data := map[string]string{
		"username":     "admin",
		"password":     "0192023a7bbd73250516f069df18b500",
		"rememberName": "true",
		"execution":    "ca308d5a-807d-443e-ac1e-e99b347ea88a_ZXlKaGJHY2lPaUpJVXpVeE1pSjkuUDVEMHNVR2IwcndKMWpaa2orUGEzZFdUR29PaFMrYXZzeWMvZ0NHcWI1ZTEzSStzQlVmcUc4Yjh5aXZQRnB5QVRnc0pNSTJRN3FmckVKdllDRlFpS3VYSHRKTHN3dkxBbjJpaGh5Zy9BVE9BTkFrQzFkb3lLK2VzUDBBaVhrZnVnaEZSRGtCUWprWVo2emI5NTVETEZzOXkrYW5Jc0xkT0dIWVNlUjlveUxucWNGWVdTNjNDZUx5MGRad3ZMc25nQmYvK2I5TGFVN2FRcGkvYVU5WDN2d2l2SXJpNkRQSEpxdzJueU5aVS9xQ0R6T1hsSnBrUFFGcFVwL1d3dzVHTlFWeDV2Q1Y3WlBETkU3VmRhQmwzR1pvV3NIWjk0ajI1SUlnNkUrNENHWnh2cmVvcnZ6M0w5QkdHTXFhc090NU5IdHdPYWFSSkVQUENncWdCa04rTFY3bit1b3BxZHR5ZU1iRWZsZGJYS1B5Ni9vaUFIZFIxbUU2UlUwOHczUU1hQy9SdDJZMXBCaGdVdEM5V2h3S2hHOEpiUkkrbnBBYkZBNXNKbWYvT09vQy9zQjh4SW9aa1BraVpFZzgyaFN0VDBOeHIvUnRDWlF1aENiREx4WDlNYnpacWVucVFuK0x4WWdZVFY4SUlrQmJLOFN0VWhpZjNXMnltM1lNeWpyN3BqYTJmQnBKenMzTndMMyszVS8xSVNCa0p0Wm5nellSUjFZRWtzNHlXUTZ6NEJ3eG5hSmhRblpqanI3VzVQMHZtaklPUWtFcTVFUS8xWWpDVVJxRndETDVubjJudUk5U1FaUnJUYmRtQlpoY2Q4SXpMb0htMS9yN1EvYnlsc0dhNWR2ODRLRmhqWGgreGk5WE1icWplaEFNSENzandUR1Y4M1B2RDhadllsd3NVRjlhRGpIeVJZSjFqTVN5QXFuWm5KYTNISU1WRzViZUFHTVljWjNOSnJZaldmNzU2bTZmT0xMYmFEc3JjQWF6RnFCN28yZUtnTmVoditYZnFxamtBYk9IeFRjMjVqZVUvNG5qY01JVklYQmFibDd2RmZleTZMUHlGVTJLMk8ycENTeDhUaFd6WExmUnAzZ2hPVjJqVitienBlQlFnN0dPeGRVYy9OZmhSWklKWEJEWUZjSmZvVk5wUDFOVUNQbEludUN4MWJ5Rmlpa0NoS3UrOGEwT1dPTTdmNmErL2dTbnZEbkVVaGhlN2ZQbDViVEZLLzA3UE1Dd09YalZ0dUw4UWdOZGhYTVVlUy9IU1JqQW95UWc5cFBDR2FYOTh3WEx6bVVvNUJ3Smk2T3pHUlVkQVNaRUdNVjRDeE42MFljbnhGOUJKMkkybmxDL2hSNm9WMVdPV1B3cjViR3RJYlc3SUMzbXFTTENQL0NyMnd2bHV3RzRKdDU0K3F2NjJNeDZuRkU5VUorcXBWck5EVGs0SythbmVxekFrc3NZd2ZHR1h5NEpwM0x2VDdOMVBzSFdDUFpQWG1CYTY0OXkzamJ2YmMxdkxxMW9RQ2pDNVpRWDZRaE9uOUJnc2xNTzl0WExoVHNQS0lleS9yVDhZamZSOFAwTlRYQzMzMktSZGJQN0tmTnI1a2ljbDYyNGVRZ1pyV3lMKzU4UEVHLzZjS25Yb2dRUG5BMHdlYWFFaG9MRDdacnF3Slc1ZHE3eGoyYnFzcFhIWkp2NDF3d3JIaVgwc09RSUJsTUp1bFNwVHpsZ05mQUM2RnlWYWVZeFBKZkJTMENScHozNTV6d1pkQlNZa3J5TEVISGV6YmdLUXJmbkNDcVJQWXNZcDFxRFRHQTBZR1F4Qk16d29OaGI1WVVNS0FQdHFoTUQ2U0NGL2FrNE9yKzEvNjU5MTRTRnAzZTV6ZXBhanV2MmxKR2I0ek5YOFRZWEErbGtlcEMzK2dScGswbVhWTGFvWmpZQnVMcHZUTm15cWkxeC9FS0xKWXRVeW5VeU5SczFvZ0RQanZiVkZmSTJ6U0x2WXUrMm9yWWhXbm5ySjRxS1FsTTRtMk1uVUY5RHMvRjlXeW40Z3Z4cEdwQXJLS2c2TmtSNktjdno1YlJNU3gvYThYM2RlNEpkRDMwUTlKNHZlRXBucHJxcjZXVGJpNVdzNC9WNTMxS1RWRkR5ZGRxL0NjRzJGTXZuTnBWY3dEMWxZVU5PMHl5Rmo4d00xb3dROHN6VE5LV3pWYlZTQ1dxanNGQUVmQXIvRzJ1OWl1NUIycFFCWmxiSi9pQ1pFVEROTkhXYm9peGNvN3RMaDRZVDZ1UDRsZW9mdSsvZXlaWWNyd0hkY0tMdkxyaDlUeXNFRk1wNTRLMTdnZGJScXVhd1ZUczVuM2p1SXZjYWoyR3NSNHJiUFNjSlFsK09LbU1UK3VVTDhnZHFyTGN2ZytESlRYeG1oOGloaGZ0RnlWTFlqZnkzKzhUQzFFTzdGL1haUEY3ZFgzalpEaGRxc0MxNnkxUkdIYUhWY0VKZ1hjKzJONk9KbkJjUE8yekNyYjNiR0JCaWl0MzN6bW8wYzhZZjUrR0hGa2l3blFIS2JWRU5vZGJZeGttaEhWT2JEK2VBM2xTaEQ0eVltREVsRnRWRk5GWDZFOFd6WmVyK0NKODhFaVNMR3RhNDgrLzJOTkJpOGh5UXBGRDB1YlNWTHBXUVV0YWJ6RFh2ckRuVnRHMkhOQ3BNVWtvN0VtK0h6UXA0R3c4cWwzVlN4b1Y1cXhFMTUxQ3NLWERMei81TGdtTnFLRXBxaENpOXhidGp5VXNqTjFLNFdpUHd2aFBqUnlRZ2xybHZXTTFUWExkYnVValh5ekxxZFVxQ0h4WW1oZ1BxVm1uY3VUODMyVk5pMm9SeDU5ek9QRGFzbU1vZllzNnBwNHY2NXYzMVVxM1BPdlYzMEZWU0tKTEV0bkdVVGI5YWh3UGZiLzErMVplb3BWbFRUQjR3Vnc0cnJaUFNJZktqS1kvRG1sL0pDR294enFKc0lud2F3bFBsblNOdjJ1UlZ3Q3lUQkw5UmdHYStkV20vWlJnYlAyU1F3NUNKbXliQVlCOGVFZHlZUyswbEFuaDhIY0tlSnhWT00zYTVOM011Yng2MzRTQjVocDlMMVpCMEEwSXd2cS8yaXVXd0ZvREo0enBOeG9hOGRlcTJGaXdIUUZBNFV4MGd5bGJXK1VkenRoZUpKbXBKZjhiRUs1TU9tY3dGdE5NUkY1ZHY5UzRqcDVhT0YwSS9QbHF4YXpOeWo2bm5xRzlvRmFqc1lsTnFta2QvbWhWalE0dTlORm10WnBCRXBieEZRWk8xNGZxR2FJbWMxWU9FY0hSYU94N1phcHBPU015NlkrcG5wSmw4UGwvdnhhOHZiZTZERDBwdlN6R0JCU2FtYjR2V1NSbDFONmd5Q0MxNFNxZFcreWIwRjRvbGY1U3dRcXBiR0dEQTgxVlU1ejVUZVFmakR6UXVTRUc2UnhpMjJVakNaN08zR1FiMUVFWTNBeTVzeXVqZjM5Q0RaMkNyRExnVTM2UWJwU2ozbGdZN2lPbmhpeW9IMEtjMTBRSGVoOTZScDdMei9CUWFYQXZnYi9PSjEvbDBMZnhrWSt6ZlYyVXFKKzdJRCtuZWxXSiticjNrM1cwdGZodDM5S2tpSTlZZWNCNUNPdWY0NzdpaTljSERNemZLTTUrTm8xS3BscEluZkVzUDZIdEFDVDlVdEtIWXg0WTBwakw4ODRhRnFEN1Rhc21CNHdINDVFdldqUHcyWWJ4a0NHSGE0T3hqQVlpN3BsTXZhTTVZRG9qNUttY29FRE1YSU1lNDlPeXRUVTd6d2VxSUt1UXlCUGIzYzFtN1Q0RTB6aUhCdEdYSG4zZW5OTmh0SmszMU1DNHltOU9ZUU5kRkZEU1pMKzJLZDYwREw5cUkvY1Z2OVozUDdOUTB3NEFpRlNldVJmSUtBUUxsYnUzVDliSW1OYjNWM1M0V3ZlVDhZOWs4TUJWNktpK2VjMnA5ZmNCd0ZQVnFlcFNnaWhMSEM0eUZuOS9pWTN4N1d0WUxPZlpLdnpKM2RKZWRjUCsreHFrNmIwVFVqcWM5QU1BYlhwRklRUU5IN2t1M2NhQU5RVU1TKzZIWUtDVjlyYUc4eHBpMndtTEl6NUkwVFRLM21LZllhOXFrYWxpQXBqRzgyRStVOFlYMzMxcEZia241VjVwQlBBSU8wVEt4d1I3T2kvZVF3STdHN3Body85QTBaZUdYOVFOZVpwelpvSXhLcE9lQ0tZQlF1RXIzZE93Nk0yRnhqRTZJdVczUWVuTWYwY2E2Mkd6MXh4RTNjYUdjYU5GTHBaYk9XRFRxTzZBdElVVGpWU2RKTDhEY3dTVXRIRk8rd0d1T3FDdDlaeW96RlVpUzVCeEd2RFFxemFRY2VFbGRoM2EwZFc1RXl1eHdGVUhzdUJxZ0UrbWNvQ21TV3ByTTNsbHlXaGJpVUtzMENySU5jMHo5Wit0dWcvU3dSQzhxOXg1QU5HMnRKM1pZajFYOUpLbkxlYjZrcUJiWlUyamV6QjZzZkU0YTJlbWFmdXR1TGk5THRHak5CK0VmWFQvZ3laanUyTCtjd1hLU0lVWEtvV1Z6YWJrczEzdXc4RDdNVlA3UDNURi9oMFd5MWdrVEF5Q3U0Szg1dHNCd0cwbkgrRW81cXVoRzBNT0VtRkg0TzZvV3dDL21MbnBmQ0NlendkakNDcCt3bktpUngxTzVsT3dLbFBORGFSenI0c2RLQ2VNNEZGRWZOSDZqTFIzczVLaElqZnRpUHFzb1c5UWloWWI3QTVPVDh1TUFZd3hUZ0pIZlptUU1YTFM5M3cweFRuYkxWaWc3N2RtbmJUT0RiWG9CblJnTUFJY1kvMTF4Z0xnaUhTN0Y4M291RUJycmdhZGRteTBWSzdzREtTK09IYXpXUUVaRk1sKzdNZzJ6S00waXRhR2pxZm1ITXJ2c2dKSm5qb3VDaUFEMXA1YmRKMDV3OFZPWFNad0VDZUhTS0xrQk1NNGhNWjkwd1Fva2psQlU9LkRHXy0yRWoyNVExdGEySXZtZ3NmOXpGdUhrMjZscG9CckdYRzU2VVozTkdPeTB3SmR0YUgtdG1aTVhETGpqSXhFdWFqc1EzcjZ5dlFfbmZQZVlvb0FB",
		"_eventId":     "submit",
		"submit":       "登录",
	}
	client := resty.New()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))
	resp, err := client.R().SetFormData(data).SetQueryParam("service", "http://10.170.208.212:9000/cas/validate").SetQueryParam("sn", "undefined").
		Post("http://10.170.208.212:40087/sso/login")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Header())
	fmt.Println(resp.String())
	fmt.Println(resp.Cookies())
	fmt.Println(resp.StatusCode())
	fmt.Println("-------")

	// client1 := resty.New()
	// client1.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))
	// client1.Header = resp.RawResponse.Header
	// client1.Cookies = resp.RawResponse.Cookies()
	// client1.Cookies = client.Cookies

	resp, err = client.R().Get("http://10.170.208.212:9000/api/uss/v1/ams/alarm/type/page?pageNo=1&pageSize=99999")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	fmt.Println("-------")
	fmt.Println(resp.Header())
	fmt.Println(client.Cookies)
	fmt.Println(resp.Cookies())
	fmt.Println(resp.RawResponse.Cookies())
	fmt.Println("-------")
}

func TestHttpOcsClient(t *testing.T) {

	client := resty.New()
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))
	resp, err := client.R().
		Post("http://10.0.133.2:8080/ocs/secret")

	if err != nil {
		fmt.Println(err)
	}

	type RetSecret struct {
		Ret string
	}
	r := RetSecret{}
	err = json.Unmarshal(resp.Body(), &r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)

	password := "123456"
	// block, _ := pem.Decode([]byte(r.Ret))
	// fmt.Println("block", block)
	// pubInterface, err:= x509.ParsePKIXPublicKey(block.Bytes)
	// pubInterface, err:= x509.ParsePKIXPublicKey([]byte(r.Ret))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	//
	// public, ok := pubInterface.(rsa.PublicKey)
	// if !ok{
	// 	fmt.Println("not ok")
	// }
	// message := []byte(password)
	// out, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &public, message, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	key := "-----BEGIN Public key-----\n" + r.Ret + "\n-----END Public key-----"
	out, err := gorsa.PublicEncrypt(password, key)
	if err != nil {
		fmt.Println(err)
	}
	data := map[string]string{
		"username": "admin",
		"password": string(out),
	}
	resp, err = client.R().SetFormData(data).Post("http://10.0.133.2:8080/ocs/login")
	if err != nil {
		fmt.Println(err)
		return
	}

	body := map[string]string{
		"areaId":    "4",
		"doorId":    "65",
		"beginDate": "2021-09-01 00:00:00 至 2021-09-08 23:59:59",
		"pageSize":  "10",
		"curPage":   "1",
	}
	resp, err = client.R().SetFormData(body).Post("http://10.0.133.2:8080/ocs/mj/doorstatusreport/datalist")
	fmt.Println(resp)
}

type ListParam struct {
	Keyword  string `json:"keyword,omitempty"`
	Order    string `json:"order,omitempty"`
	OrderBy  string `json:"orderBy,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

// 找一个或多个所属公司下的所有有用户的请求和响应
type CompanyUserListRequest struct {
	*ListParam
	CompanyNames []string `json:"companyNames,omitempty"`
	ProjectId    string   `json:"projectId"`
	RequestId    string   `json:"-"`
}

type User struct {
	Name              string   `json:"name"`
	LoginType         string   `json:"loginType"`
	UserId            string   `json:"userId"`
	Email             string   `json:"email"`
	PhoneNumber       string   `json:"phoneNumber"`
	State             string   `json:"state"`
	Status            int      `json:"status"`
	Admin             bool     `json:"admin"`
	OrgPermission     string   `json:"orgPermission"`
	ProjectPermission string   `json:"projectPermission"`
	OrganizationID    string   `json:"organizationId"`
	OrganizationName  string   `json:"organizationName"`
	CompanyName       string   `json:"companyName"`
	CreateTime        string   `json:"createTime"`
	Creator           string   `json:"creator"`
	PasswordState     string   `json:"passwordState"`
	PasswordSetted    bool     `json:"PasswordSetted"`
	NeedResetPassword bool     `json:"needResetPassword"`
	RoleNameList      []string `json:"roleNameList"`
	RoleIds           []string `json:"roleIds"`
	Password          string   `json:"password"`
	AccessKey         string   `json:"accessKey"`
	SecretKey         string   `json:"secretKey"`
}

type UserListResponse struct {
	TotalCount int    `json:"totalCount"`
	Result     []User `json:"result"`
}

func TestHttpIdaasClient(t *testing.T) {
	var userList = UserListResponse{}
	request := CompanyUserListRequest{
		ListParam: &ListParam{
			Keyword:  "",
			Order:    "DESC",
			OrderBy:  "create_time",
			PageNo:   1,
			PageSize: 10,
		},
		CompanyNames: []string{"招商局3"},
		ProjectId:    "148dbb10005e11ea939c7df8991ad0e7",
		RequestId:    "",
	}

	client := http.Client{}

	data, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
	}
	body := bytes.NewBuffer(data)
	url := "http://localhost:8667/v1/company/UserLists"
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err)
	}

	//req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "idaas-csrftoken=t91xgvjqzyt21kg3; idaas-default-url=; idaas-project-name=deng; idaas-sessionid=pbbty6ihprc3qe0x")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(respData, &userList)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", userList)
}
