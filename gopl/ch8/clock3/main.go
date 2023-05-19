package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

// ===================没有并发的clock服务====================
var PortMap = map[int]string{
	8000: "Beijing",
	8010: "Network",
	8020: "Tokyo",
	8030: "London",
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "端口号")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		conn, err := listener.Accept() // 阻塞，直到有新的客户端连接
		if err != nil {
			log.Print(err.Error())
			continue
		}
		go handleConn(conn, port) // 服务器程序同一时间只能处理一个客户端连接
	}
}

func handleConn(c net.Conn, port int) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, fmt.Sprintf("%s time %s", PortMap[port], time.Now().Format("15:04:05\n")))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
