package constant_test

import (
	"fmt"
	"testing"
)

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

func TestConstantTry(t *testing.T) {
	fmt.Println(Monday, Tuesday)
}

func TestContantTry1(t *testing.T) {
	a := 1 // 0111
	fmt.Println(Readable, Writeable, Executable)
	fmt.Println(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
