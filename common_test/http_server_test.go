package common_test

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/gin-contrib/location"
)

func HttpForward(w http.ResponseWriter, req *http.Request) {
	// req.host不带http,req.url是完整的url
	// fmt.Println(req.Host, " ", req.URL, "\n")

	fmt.Println("url: ", req.URL)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Print("io.ReadFull(req.Body, body) ", err.Error())
		// return,没有数据也是可以的，不需要直接结束
	}
	fmt.Print("req count :", len(body), "\n")

	reqUrl := req.URL.String()

	fmt.Printf("reqUrl : %s, body: %s", reqUrl, string(body))
	cli := &http.Client{}
	reqForward, err := http.NewRequest(req.Method, reqUrl, strings.NewReader(string(body)))
	if err != nil {
		fmt.Print("http.NewRequest ", err.Error())
		return
	}

	for k, v := range reqForward.Header {
		reqForward.Header.Set(k, v[0])
	}
	res, err := cli.Do(reqForward)
	if err != nil {
		fmt.Print("cli.Do(reqForward) ", err.Error())
		return
	}
	defer res.Body.Close()

	for k, v := range res.Header {
		w.Header().Set(k, v[0])
	}
	_, err = io.Copy(w, res.Body)
	if err != nil {
		fmt.Print("io.Copy ", err.Error())
		return
	}
}

func TestHttpServer(t *testing.T) {
	http.HandleFunc("/", HttpForward)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func EventHandle(ctx *gin.Context) {
	fmt.Println("url: ", ctx.Request.URL)

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Print("io.ReadFull(req.Body, body) ", err.Error())
		// return,没有数据也是可以的，不需要直接结束
	}
	fmt.Print("req count :", len(body), "\n")

	ctx.JSON(200, "success")
}

func TestHttpServe1r(t *testing.T) {
	gin.SetMode(gin.DebugMode)

	gin.ForceConsoleColor()

	gin.DefaultWriter = logrus.StandardLogger().Out

	gin.DefaultErrorWriter = logrus.StandardLogger().Out

	server := gin.Default()

	server.Use(location.Default())

	server.Any("/", EventHandle)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: server,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Errorf("listen: %s\n", err.Error())
		}
	}()

	time.Sleep(5* time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		logrus.Println("timeout of 5 seconds.")
	}
	logrus.Println("Server exiting")

	select {

	}
}