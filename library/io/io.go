package io

import "errors"

const (
	SeekStart   = 0
	SeekCurrent = 1
	SeekEnd     = 2
)

var ErrShortWrite = errors.New("short write")

var errInvalidaWrite = errors.New("invalid write result")
