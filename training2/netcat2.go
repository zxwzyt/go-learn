// package main

// import (
// 	"log"
// 	"net"
// 	"os"
// 	"zxw-go/common"
// )

// func main() {
// 	conn, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer conn.Close()
// 	go common.MustCopy(os.Stdout, conn)
// 	common.MustCopy(conn, os.Stdout)
// }
