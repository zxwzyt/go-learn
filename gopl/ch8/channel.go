// package main

// import (
// 	"go-learn/utils"
// 	"io"
// 	"log"
// 	"net"
// 	"os"
// )

// // 不太明白
// func main() {
// 	// ch := make(chan int) // unbuffered channel
// 	conn, err := net.Dial("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	done := make(chan struct{})
// 	go func() {
// 		io.Copy(os.Stdout, conn)
// 		log.Println("done")
// 		done <- struct{}{}
// 	}()
// 	utils.MustCopy(conn, os.Stdin)
// 	conn.Close()
// 	<-done
// }
