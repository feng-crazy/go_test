package utils

import (
	"errors"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	RPCREQUEST_TIMEOUT = 10
	RECONNECT_TIMEOUT  = 2
)

type RpcClient struct {
	mutex     sync.Mutex
	rpcClient *rpc.Client
	conn      net.Conn
	address   string
}

func (c *RpcClient) Connect(address string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.connect(address)
}

func (c *RpcClient) connect(address string) error {
	c.address = address
	count := 0
	for {
		conn, err := net.DialTimeout("tcp", address, time.Second*RECONNECT_TIMEOUT)
		if err != nil {
			log.Info("Rpc Client try to Connect ", address)
			time.Sleep(500 * time.Millisecond)
			count++
			if count >= 2 {
				return errors.New("Dial Timeout")
			}
			continue
		}

		c.conn = conn
		c.rpcClient = jsonrpc.NewClient(c.conn)
		break
	}
	return nil
}

func (c *RpcClient) DisConnect() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.rpcClient != nil {
		c.rpcClient.Close()
		c.rpcClient = nil
	}
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}

	return nil
}

func (c *RpcClient) Call(Method string, Req interface{}, Res interface{}, timeout int) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if timeout <= 0 || timeout > 20 {
		timeout = RPCREQUEST_TIMEOUT
	}

	if c.conn == nil || c.rpcClient == nil {
		if err := c.connect(c.address); err != nil {
			return err
		}
	}

	defer c.conn.SetReadDeadline(time.Time{})
	if err := c.conn.SetReadDeadline(time.Now().Add(time.Duration(timeout) * time.Second)); err != nil {
		return err
	}
	if err := c.rpcClient.Call(Method, Req, Res); err != nil {
		if strings.Contains(err.Error(), "connection is shut down") {
			if err = c.connect(c.address); err != nil {
				return err
			}
			return c.rpcClient.Call(Method, Req, Res)
		}
		return err
	}
	return nil
}

// ////////////////////////////////////////////////////////////////

func NewRpcClient() *RpcClient {
	return &RpcClient{}
}
