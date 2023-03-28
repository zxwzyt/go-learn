package training

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// 	"time"
// )

// func main() {
// 	listener, err := net.Listen("tcp", "localhost:8000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for {
// 		conn, err := listener.Accept() // 阻塞，直到有新的客户端连接
// 		if err != nil {
// 			log.Print(err)
// 			continue
// 		}
// 		fmt.Printf("%s", "test")
// 		handleConn(conn) // 服务器程序同一时间只能处理一个客户端连接
// 	}
// }

// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
// 		if err != nil {
// 			return
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }
