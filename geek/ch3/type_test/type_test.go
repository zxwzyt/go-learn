package type_test

import (
	"fmt"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64 = 2
	// b = a 不支持隐式类型转换, 要用显示类型转换
	var c MyInt
	// b = c 别名隐式转换也不可以
	fmt.Println(a, b, c)
}
