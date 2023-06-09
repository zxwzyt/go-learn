package bytes

import (
	"go-learn/library/internal/bytealg"
	"unicode/utf8"
)

func Equal(a, b []byte) bool {
	return string(a) == string(b)
}

func Compare(a, b []byte) int {
	return bytealg.Compare(a, b)
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

// func Contains(b, subslice []byte) bool {
// 	return Index(b, subslice) != -1
// }

// func ContainsAny(b []byte, chars string) bool {
// 	return IndexAny(b, chars) >= 0
// }

// func ContainsRune(b []byte, r rune) bool {
// 	return IndexRune(b, r) >= 0
// }

// func indexBytePortable(s []byte, c byte) int {
// 	for i, b := range s {
// 		if b == c {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func LastIndex(s, sep []byte) int {
// 	n := len(sep)
// 	switch {
// 	case n == 0:
// 		return len(s)
// 	case n == 1:
// 		return LastIndexByte(s, sep[0])
// 	case n == len(s):
// 		if Equal(s, sep) {
// 			return 0
// 		}
// 		return -1
// 	case n > len(s):
// 		return -1
// 	}
// 	hashss, pow := bytealg.HashStrBytes(sep)
// 	last := len(s) - n
// 	var h uint32
// 	for i := len(s) - 1; i >= last; i-- {
// 		h = h*bytealg.PrimeRK + uint32(s[i])
// 	}
// 	if h == hashss && Equal(s[last:], sep) {
// 		return last
// 	}
// 	for i := last - 1; i >= 0; i-- {
// 		h *= bytealg.PrimeRK
// 		h += uint32(s[i])
// 		h -= pow * uint32(s[i+n])
// 		if h == hashss && Equal(s[i:i+n], sep) {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func IndexAny(s []byte, chars string) int {
// 	if chars == "" {
// 		return -1
// 	}
// 	if len(s) == 1 {
// 		r := rune(s[0])
// 		if r >= utf8.RuneSelf {
// 			for _, r = range chars {
// 				if r == utf8.RuneError {
// 					return 0
// 				}
// 			}
// 		}
// 		if bytealg.IndexByteString(chars, s[0]) >= 0 {
// 			return 0
// 		}
// 		return -1
// 	}

// 	if len(chars) == 1 {
// 		r := rune(chars[0])
// 		if r >= utf8.RuneSelf {
// 			r = utf8.RuneError
// 		}
// 		return IndexRune(s, r)
// 	}

// 	if len(s) > 8 {
// 		if as, isASCII := makeASCIISet(chars); isASCII {
// 			for i, c := range s {
// 				if as.contains(c) {
// 					return i
// 				}
// 			}
// 			return -1
// 		}
// 	}
// 	var width int
// 	for i := 0; i < len(s); i += width {
// 		r := rune(s[i])
// 		if r < utf8.RuneSelf {
// 			if bytealg.IndexByteString(chars, s[i]) >= 0 {
// 				return i
// 			}
// 			width = 1
// 			continue
// 		}

// 		r, width = utf8.DecodeRune(s[i:])
// 		if r != utf8.RuneError {
// 			if len(chars) == width {
// 				if chars == string(r) {
// 					return i
// 				}
// 				continue
// 			}

// 			if bytealg.MaxLen >= width {
// 				if bytealg.IndexString(chars, string(r)) >= 0 {
// 					return i
// 				}
// 				continue
// 			}
// 		}
// 		for _, ch := range chars {
// 			if r == ch {
// 				return i
// 			}
// 		}
// 	}

// 	return -1
// }

// type asciiSet [8]uint32

// func makeASCIISet(chars string) (as asciiSet, ok bool) {
// 	for i := 0; i < len(chars); i++ {
// 		c := chars[i]
// 		if c >= utf8.RuneSelf {
// 			return as, false
// 		}
// 		as[c>>5] |= 1 << uint(c&31)
// 	}
// 	return as, true
// }

// func LastIndexByte(s []byte, c byte) int {
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if s[i] == c {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func IndexRune(s []byte, r rune) int {
// 	return 0
// }

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

// // contains reports whether c is inside the set.
// func (as *asciiSet) contains(c byte) bool {
// 	return (as[c>>5] & (1 << uint(c&31))) != 0
// }

func main() {

}
