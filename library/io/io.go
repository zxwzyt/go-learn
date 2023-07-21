// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package io provides basic interfaces to I/O primitives.
// Its primary job is to wrap existing implementations of such primitives.
// such as those in package os, into shared public interfaces that
// abstract the functionality, plus some other related primitives.

// Because these interfaces and primitivies warp lower-level operations with
// various implementations, unless otherwise informed clients shoud not
// assume they are safe for parallel execution.

package io

import "errors"

const (
	SeekStart   = 0
	SeekCurrent = 1
	SeekEnd     = 2
)

var ErrShortWrite = errors.New("short write")

var errInvalidWrite = errors.New("invalid write result")

var ErrShortBuffer = errors.New("short buffer")

var EOF = errors.New("EOF")

var ErrUnexpectedEOF = errors.New("unexpected EOF")

var ErrNoProgress = errors.New("multiple Read calls return no data or error")

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadCloser struct {
	Reader
	Closer
}

type WriteCloser struct {
	Writer
	Closer
}

type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

type ReadSeeker interface {
	Reader
	Seeker
}

type ReadSeekerCloser interface {
	Reader
	Seeker
	Closer
}

type WriteSeeker interface {
	Writer
	Seeker
}

type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}

type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}

type WriterTo interface {
	WriteTo(w Writer) (n int64, err error)
}

type ReadAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}

type ByteReader interface {
	ReadByte() (byte, error)
}

type BytesScanner interface {
	ByteReader
	UnreadByte() error
}

type ByteWriter interface {
	WriteByte(c byte) error
}

type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}

type StringWriter interface {
	WriteString(s string) (n int, err error)
}

func WriteString(w Writer, s string) (n int, err error) {
	if sw, ok := w.(StringWriter); ok {
		return sw.WriteString(s)
	}
	return w.Write([]byte(s))
}

func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error) {
	if len(buf) < min {
		return 0, ErrShortBuffer
	}
	for n < min && err == nil {
		var nn int
		nn, err = r.Read(buf[n:])
		n += nn
	}
	if n >= min {
		err = nil
	} else if n > 0 && err == EOF {
		err = ErrUnexpectedEOF
	}
	return
}

func ReadFull(r Reader, buf []byte) (n int, err error) {
	return ReadAtLeast(r, buf, len(buf))
}

func CopyN(dst Writer, src Reader, n int64) (written int64, err error) {
	written, err = Copy(dst, LimitReader(src, n))
	if written == n {
		return n, nil
	}
	if written < n && err == nil {
		err = EOF
	}
	return
}

func Copy(dst Writer, src Reader) (written int64, err error) {
	return copyBuffer(dst, src, nil)
}

func copyBuffer(dst Writer, src Reader, buf []byte) (writen int64, err error) {
	return
}
