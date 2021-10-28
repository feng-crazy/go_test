package gorilla_test

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sacOO7/gowebsocket"
	"github.com/sirupsen/logrus"
)

func TestGorillaClient(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	addr := "10.0.128.5:14305"

	u := url.URL{Scheme: "ws", Host: addr, Path: "/opensmart/api/1.0/login?user=system"}
	logrus.Printf("connecting to %s", u.String())

	// websocket.Dialer{
	// 	NetDial:           nil,
	// 	NetDialContext:    nil,
	// 	Proxy:             nil,
	// 	TLSClientConfig:   nil,
	// 	HandshakeTimeout:  0,
	// 	ReadBufferSize:    0,
	// 	WriteBufferSize:   0,
	// 	WriteBufferPool:   nil,
	// 	Subprotocols:      nil,
	// 	EnableCompression: false,
	// 	Jar:               nil,
	// }
	logrus.Println(u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		logrus.Fatal("dial:", err)
		return
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				logrus.Println("read:", err)
				return
			}
			logrus.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				logrus.Println("write:", err)
				return
			}
			logrus.Println("write success:", t.String())
		case <-interrupt:
			logrus.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logrus.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

type ConnectionOptions struct {
	UseCompression bool
	UseSSL         bool
	Proxy          func(*http.Request) (*url.URL, error)
	Subprotocols   [] string
}

func TestGorillaClient1(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	WebsocketDialer := &websocket.Dialer{}
	RequestHeader:= http.Header{}

	RequestHeader.Set("Accept-Encoding", "gzip, deflate, sdch")
	RequestHeader.Set("Accept-Language", "en-US,en;q=0.8")
	RequestHeader.Set("Pragma", "no-cache")
	RequestHeader.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36")

	ConnectionOptions := ConnectionOptions{
		UseCompression: false,
		UseSSL:         true,
	}

	WebsocketDialer.EnableCompression = ConnectionOptions.UseCompression
	WebsocketDialer.TLSClientConfig = &tls.Config{InsecureSkipVerify: ConnectionOptions.UseSSL}
	WebsocketDialer.Proxy = ConnectionOptions.Proxy
	WebsocketDialer.Subprotocols = ConnectionOptions.Subprotocols

	addr := "10.0.128.5:14305"
	u := url.URL{Scheme: "ws", Host: addr, Path: "/opensmart/api/1.0/login?user=system"}
	logrus.Printf("connecting to %s", u.String())

	c, resp, err := WebsocketDialer.Dial(u.String(), RequestHeader)
	res, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(res))
	if err != nil {
		logrus.Fatal("dial:", err)
		return
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				logrus.Println("read:", err)
				return
			}
			logrus.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				logrus.Println("write:", err)
				return
			}
			logrus.Println("write success:", t.String())
		case <-interrupt:
			logrus.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				logrus.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func TestClient(t *testing.T) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)


	addr := "10.0.128.5:14305"

	u := url.URL{Scheme: "ws", Host: addr, Path: "/opensmart/api/1.0/login?user=system"}
	logrus.Printf("connecting to %s", u.String())

	// socket := gowebsocket.New("ws://10.0.128.5:14305/opensmart/api/1.0/login?user=system")
	socket := gowebsocket.New("ws://10.0.128.5:14305/opensmart/api/1.0/login%3Fuser=system")
	// socket := gowebsocket.New(u.String())

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Println("Recieved connect error ", err)
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		log.Println("Recieved message " + message)
	}

	socket.OnBinaryMessage = func(data [] byte, socket gowebsocket.Socket) {
		log.Println("Recieved binary data ", data)
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved ping " + data)
	}

	socket.OnPongReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved pong " + data)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	socket.Connect()

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return
		}
	}
}