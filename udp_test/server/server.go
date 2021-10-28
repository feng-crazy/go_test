package main

import (
	"encoding/binary"
	"flag"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "6002", "port")

func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		logrus.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		logrus.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var err error
	err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		logrus.Println(err)
		return
	}
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		logrus.Println("failed to read UDP msg because of ", err.Error())
		return
	}

	logrus.Println(n, remoteAddr)
	logrus.Println("data:", string(data[:n]))

	b := make([]byte, 4)
	daytime := time.Now().Unix()
	binary.BigEndian.PutUint32(b, uint32(daytime))
	n, err = conn.WriteToUDP(b, remoteAddr)
	if err != nil {
		logrus.Println("failed to write UDP msg because of ", err.Error())
		return
	}

}
