package ch8

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Hello World")
	// os.Exit(-1) // 异常退出
	// fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("Hello World", os.Args[1])
	}
}
