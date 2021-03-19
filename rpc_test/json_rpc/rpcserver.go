package utils

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"

	log "github.com/sirupsen/logrus"
)

type RpcServer struct {
	Flag      bool
	wg        sync.WaitGroup
	listener  net.Listener
	rpcServer *rpc.Server
}

func (rs *RpcServer) Init(addr string, handler ...interface{}) error {
	if !rs.Flag {
		rs.rpcServer = rpc.NewServer()
		for _, h := range handler {
			rs.rpcServer.Register(h)
		}

		tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
		if err != nil {
			log.Error(err)
		}
		listener, err := net.ListenTCP("tcp", tcpAddr)
		rs.listener = listener
		if err != nil {
			log.Error("Rpc Listen Error:", err)
			return err
		}
		rs.wg.Add(1)
		go func() {
			for {
				conn, err := rs.listener.Accept()
				if err != nil {
					log.Error("Rpc Server Accept:", err)
					// continue
					break
				}
				go rs.rpcServer.ServeCodec(jsonrpc.NewServerCodec(conn))
			}
			rs.wg.Done()
		}()
		rs.Flag = true
	}
	return nil
}

func (rs *RpcServer) DeInit() error {
	if rs.Flag {
		rs.listener.Close()
		rs.wg.Wait()
		rs.Flag = false
	}
	return nil
}

// ////////////////////////////////////////////////////////
func NewRpcServer() *RpcServer {
	return &RpcServer{}
}
