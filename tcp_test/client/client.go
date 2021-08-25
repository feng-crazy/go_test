package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"go_test/tcp_test/protocol"
)

func main() {
	for {
		server := "127.0.0.1:9988"
		tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
			time.Sleep(3 * time.Second)
			continue
		}

		fmt.Println("connect success")
		buffer := make([]byte, 1024)
		for {
			words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
			n, err := conn.Write(protocol.Packet([]byte(words)))
			fmt.Println(n, err)
			if err != nil {
				fmt.Println(conn.RemoteAddr().String(), "Write connection error: ", err)
				break
			}

			n, err = conn.Read(buffer)
			if err != nil {
				fmt.Println(conn.RemoteAddr().String(), "read connection error: ", err)
				break
			}
			fmt.Println("读数据应答", string(buffer[:n]))
			time.Sleep(1 * 1e9)
		}

		_ = conn.Close()
	}

}
