package bytes

import (
	"internal/bytealg"
	"unicode/utf8"
)

func Equal(a, b []byte) bool {
	return string(a) == string(b)
}

func Compare(a, b []byte) int {
	return bytealg.Compare(a, b)
	return 0
}

func explodes(s []byte, n int) [][]byte {
	if n <= 0 {
		n = len(s)
	}

	a := make([][]byte, n)
	var size int
	na := 0
	for len(s) > 0 {
		if na+1 >= n {
			a[na] = s
			na++
			break
		}
		_, size = utf8.DecodeRune(s)
		a[na] = s[0:size:size]
		s = s[size:]
		na++
	}
	return a[0:na]
}

func Count(s, sep []byte) int {
	if len(sep) == 0 {
		return utf8.RuneCount(s) + 1
	}
	if len(sep) == 1 {
		return bytealg.Count(s, sep[0])
	}
	n := 0
	for {
		i := Index(s, sep)
		if i == -1 {
			return n
		}
		n++
		s = s[i+len(sep):]
	}
}

func Contains(b, subslice []byte) bool {
	return Index(b, subslice) != -1
}

func ContainsAny(b []byte, chars string) bool {
	return IndexAny(b, chars) >= 0
}

func ContainsRune(b []byte, r rune) bool {
	return IndexRune(b, r) >= 0
}

func IndexAny(s []byte, chars string) int {
	return 0
}

func IndexRune(s []byte, r rune) int {
	return 0
}

func Index(s, sep []byte) int {
	n := len(sep)
	switch {
	case n == 0:
		return 0
	case n == 1:
		return IndexByte(s, sep[0])
	case n == len(s):
		if Equal(sep, s) {
			return 0
		}
		return -1
	case n > len(s):
		return -1
	case n <= bytealg.MaxLen:
		if len(s) <= bytealg.MaxBruteForce {
			return bytealg.Index(s, sep)
		}

		c0 := sep[0]
		c1 := sep[1]
		i := 0
		t := len(s) - n + 1
		fails := 0
		for i < t {
			if s[i] != c0 {
				o := IndexByte(s[i+1:t], c0)
				if o < 0 {
					return -1
				}
				i += o + 1
			}
			if s[i+1] == c1 && Equal(s[i:i+n], sep) {
				return i
			}
			fails++
			i++
			if fails > bytealg.Cutover(i) {
				r := bytealg.Index(s[i:], sep)
				if r >= 0 {
					return r + i
				}
				return -1
			}
		}
		return -1
	}

	c0 := sep[0]
	c1 := sep[1]
	i := 0
	fails := 0
	t := len(s) - n + 1
	for i < t {
		if s[i] != c0 {
			o := IndexByte(s[i+1:t], c0)
			if o < 0 {
				break
			}
			i += o + 1
		}
		if s[i+1] == c1 && Equal(s[i:i+n], sep) {
			return i
		}
		i++
		fails++
		if fails >= 4+i>>4 && i < t {
			j := bytealg.IndexRabinKarpBytes(s[i:], sep)
			if j < 0 {
				return -1
			}
			return i + j
		}
	}
	return -1
}

func IndexByte(b []byte, c byte) int {
	return bytealg.IndexByte(b, c)
}
