package gin_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func ReverseProxy() gin.HandlerFunc {
	target := "localhost:9081"
	return func(c *gin.Context) {
		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = target
			req.URL.Path = c.Param("url")
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func TestProxy(t *testing.T) {
	// 创建默认的引擎
	r := gin.Default()

	r.Any("/kuiper/*url", ReverseProxy())
	// 运行 默认为80端口
	r.Run()
}

func TestProxyExample(t *testing.T) {
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this call was relayed by the reverse proxy")
	}))
	defer backendServer.Close()

	rpURL, err := url.Parse(backendServer.URL)
	if err != nil {
		log.Fatal(err)
	}
	frontendProxy := httptest.NewServer(httputil.NewSingleHostReverseProxy(rpURL))
	defer frontendProxy.Close()

	resp, err := http.Get(frontendProxy.URL)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

}
