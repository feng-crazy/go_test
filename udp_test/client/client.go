package main

import (
	"encoding/binary"
	"flag"
	"net"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "6002", "port")

func main() {
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		logrus.Println("Can't resolve address: ", err)
		time.Sleep(2 * time.Second)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		logrus.Println("Can't dial: ", err)
		time.Sleep(2 * time.Second)
	}

	go func() {
		for {
			data := make([]byte, 1024)
			n, err := conn.Read(data)
			if err != nil {
				logrus.Println("failed to read UDP msg because of ", err)
				continue
			}
			t := binary.BigEndian.Uint32(data[:n])
			logrus.Println(time.Unix(int64(t), 0).String())
		}
	}()

	go func() {
		i := 0
		for {
			// _, err = conn.WriteTo([]byte("tset1"), addr)
			_, err = conn.Write([]byte("tset" + strconv.Itoa(i)))
			if err != nil {
				logrus.Println("failed:", err)
				time.Sleep(2 * time.Second)
				continue
			}
			logrus.Printf("write.... %d\n", i)

			time.Sleep(2 * time.Second)

			i++
		}
	}()

	select {}
}
