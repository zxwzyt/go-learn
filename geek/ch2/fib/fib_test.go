package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {

	var (
		a int = 1
		b     = 2
	)
	// var a int = 1
	// var b int = 2
	fmt.Println(a)
	for i := 0; i < 5; i++ {
		fmt.Println("", b)
		tmp := a
		a = b
		b = tmp + a
	}
	t.Logf("%d", b)
}

// 同时赋值多个变量
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	fmt.Println(a, b)
}
