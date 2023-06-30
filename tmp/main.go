package main

func main() {
	// x := []int{1, 2, 3, 4, 5}
	// y := x[1:3]
	// x = append(x, 8)
	// y = append(y, 9)
	// y = append(y, 9)
	// y = append(y, 9)
	// fmt.Printf("x len:%d, cap:%d\r\n", len(x), cap(x))
	// fmt.Printf("y len:%d, cap:%d\r\n", len(y), cap(y))

	// fmt.Println(add.Add(2, 15))
	main()
}

// func Demo1() {
// s := make([]byte, 200)
// ptr := unsafe.Pointer(&s[0])

// // 通过指针构造slice
// var ptr1 unsafe.Pointer
// s1 := ((*[1 << 10]byte)(ptr1))[:200]

// 	var ptr2 unsafe.Pointer
// 	var s2 := struct {
// 		add uintpptr
// 		len int
// 		cap int
// 	}{ptr2, length, length}
// 	s3 := *(*[]byte)(unsafe.Pointer(&s2))

// 	var o []byte
// 	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&o))
// 	sliceHeader.Cap = length
// 	sliceHeader.Len = length
// 	sliceHeader.Data = uintptr(ptr)
// }

// struct Hmap {
// 	uint8 B
// 	uint16 bucketsize

// 	byte *buckets
// 	byte *oldbuckets
// }

// struct Bucket {
// 	uint tophash[BUCKETSIZE]
// 	Bucket *overflow
// 	byte data[1]
// }
