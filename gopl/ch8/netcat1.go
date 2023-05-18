package main

import (
	"go-learn/utils"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	utils.MustCopy(os.Stdout, conn)
}
