package training

import (
	"fmt"
)

func Pipe3() {
	// 执行以下方法会报错，因为在同一无缓存通道，在接收者为准备好之前，发送操作是阻塞的。第三行发生在第二行之后，所以发送操作执行时接收者还未准备好
	// ch := make(chan string, 3)
	// ch <- "A"
	// fmt.Println(<-ch)

	// 解决方法一，添加配对接收者，方法是接收和发送在两个goroutine中

	ch := make(chan string)
	go func() {
		ch <- "A"
	}()

	fmt.Println(<-ch)

	// 方法二，使用缓存通道,好处是解耦接收和发送的goroutine
	ch1 := make(chan string, 3)
	ch1 <- "A"
	ch1 <- "B"
	ch1 <- "C"
	fmt.Println(<-ch1)
	fmt.Println(cap(ch1))
	fmt.Println(len(ch1))
	ch1 <- "D"
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	mirroredQuery()
}

func mirroredQuery() {
	responses := make(chan string, 3)
	go func() {
		responses <- request("baidu.com")
	}()

	go func() {
		responses <- request("google.qiyi.com")
	}()

	go func() {
		responses <- request("idc.qiyi.com")
	}()

	fmt.Println(<-responses)
	fmt.Println(<-responses)
	fmt.Println(<-responses)
}

func request(hostname string) (response string) {
	return hostname
}
