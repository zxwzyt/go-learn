package main

import (
	"go-learn/utils"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	for _, v := range os.Args[1:] {
		if v != "" {
			addr := strings.Split(v, "=")
			dial(addr[1])
		}
	}
}

func dial(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go utils.MustCopy(os.Stdout, conn)
}
