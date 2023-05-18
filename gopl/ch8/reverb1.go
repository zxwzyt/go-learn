package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"net"
// 	"strings"
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
// 		go handleConn(conn) // 并发执行，不等待返回；服务端可同时处理多个客户端连接
// 	}
// }

// func echo(c net.Conn, shout string, delay time.Duration) {
// 	fmt.Fprintln(c, "\t\t", strings.ToUpper(shout), strings.ToUpper("Start"))
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t\t", shout, strings.ToUpper("middle"))
// 	time.Sleep(delay)
// 	fmt.Fprintln(c, "\t\t", strings.ToUpper(shout), strings.ToLower("end"))
// }

// func handleConn(c net.Conn) {
// 	input := bufio.NewScanner(c)
// 	for input.Scan() { // 阻塞？不太理解
// 		fmt.Fprintln(c, "\t", "Conn")
// 		fmt.Fprintln(c, "\t", input.Text())
// 		go echo(c, input.Text(), 3*time.Second) // go后跟的函数的参数会在go语句自身执行时被求值；因此input.Text()会在main goroutine中被求值。如果不用go并发，会阻塞直到前一个echo执行完
// 	}
// 	c.Close()
// }
