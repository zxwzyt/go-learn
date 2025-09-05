package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	//	for i := 0; i < 10; i++ {
	//		wg.Add(1)
	//		go hello(i)
	//	}
	//
	// wg.Wait()

	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i=%d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
