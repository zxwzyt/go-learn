package main

import (
	"io"
	"log"
	"net"
	"os"
	"zxw-go/common"
)

// 不太明白
func main() {
	// ch := make(chan int) // unbuffered channel
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	common.MustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}
