package bytes_test

import (
	"fmt"
	"go-learn/library/bytes"
	"testing"
)

func TestCount(t *testing.T) {
	s := []byte("abcdef")
	sep := []byte("c")
	ret := bytes.Count(s, sep)
	fmt.Println(ret)
}
