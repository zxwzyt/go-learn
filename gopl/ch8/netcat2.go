package main

import (
	"go-learn/utils"
	"log"
	"net"
	"os"
)

func netcat2() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	go utils.MustCopy(os.Stdout, conn)
	utils.MustCopy(conn, os.Stdout)
}
