package http_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"
)

func TestParseHtml(t *testing.T) {
	url := "https://www.sina.com.cn/mid/ask/list.shtml"
	file, err := ioutil.ReadFile("test.shtml")
	if err != nil {
		t.Error(err)
		return
	}

	data, err := Convert2Utf8(file, "UTF-8")
	if err != nil {
		t.Error(err)
		return
	}

	urlList, err := GetUrlList(data, url)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("list %+v\n", urlList)
}

func Convert2Utf8(raw []byte, contentType string) ([]byte, error) {
	reader := bytes.NewReader(raw)

	utf8Reader, err := charset.NewReader(reader, contentType)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(utf8Reader)
}

// Get all sub url of node.
func getUrlList(node *html.Node, refUrl *url.URL) []string {
	var urlList []string
	log.Printf("node.Namespace: %s node.Data: %s node.Type: %d", node.Namespace, node.Data, node.Type)
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			log.Printf("key: %s val: %s namespase: %s", a.Key, a.Val, a.Namespace)
			if a.Key == "href" {
				if a.Val != "javascript:;" && a.Val != "javascript:void(0)" {
					subUrl, err := refUrl.Parse(a.Val)
					if err == nil {
						urlList = append(urlList, subUrl.String())
					}
				}
				break
			}
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		//log.Printf("node : %+v\n", child)
		childUrlList := getUrlList(child, refUrl)
		urlList = append(urlList, childUrlList...)
	}

	return urlList
}

// GetUrlList Get sub url list of pre url from data.
func GetUrlList(data []byte, preUrl string) ([]string, error) {
	// parse html
	node, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("html.Parse(): %s", err.Error())
	}

	// parse url
	refUrl, err := url.ParseRequestURI(preUrl)
	if err != nil {
		return nil, fmt.Errorf("%s: url.ParseRequestURL(): %s", preUrl, err.Error())
	}

	urlList := getUrlList(node, refUrl)

	return urlList, nil
}

// ParseHostName Parse hostname from raw url.
func ParseHostName(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	if u.Host == "" {
		return "", fmt.Errorf("empty host")
	}

	// 可能出现如xxx.baidu.com:8080这样带端口号的情况
	hostName := strings.Split(u.Host, ":")
	if len(hostName) == 0 {
		return "", fmt.Errorf("invalid hostname")
	}

	return hostName[0], nil
}
