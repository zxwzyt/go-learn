package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})

	if err != nil {
		fmt.Println("连接服务端失败, err:", err)
		return
	}
	defer socket.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		sendData := strings.Trim(input, "\r\n")
		if strings.ToLower(sendData) == "q" {
			return
		}
		_, err = socket.Write([]byte(sendData))
		if err != nil {
			fmt.Println("发送数据失败, err:", err)
			return
		}
		data := make([]byte, 3)
		n, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			fmt.Println("接收数据失败, err:", err)
			return
		}
		fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
	}
}
