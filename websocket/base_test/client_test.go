package base_test

import (
	"net/url"
	"os"
	"os/signal"
	"testing"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
)

func TestStdWebsocketClient(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	addr := "10.0.128.5:14305"

	u := url.URL{Scheme: "ws", Host: addr, Path: "/opensmart/api/1.0/login?user=system"}
	logrus.Printf("connecting to %s", u.String())

	origin := "http://10.0.128.5:14305/opensmart/api/1.0/login?user=system"
	url := "ws://10.0.128.5:14305/opensmart/api/1.0/login?user=system"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		logrus.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		logrus.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		logrus.Fatal(err)
	}
	logrus.Printf("Received: %s.\n", msg[:n])
}
