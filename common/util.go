package common

import (
	"io"
	"log"
)

// 从连接中读取数据并将内容写入到标准输出
func MustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
