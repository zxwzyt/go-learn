package training

import "fmt"

func Pipe() {
	naturals := make(chan int)
	squares := make(chan int)
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// main goroutine
	for {
		fmt.Println(<-squares)
	}
}

func Pipe2() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// main goroutine
	for x := range squares {
		fmt.Println(x)
	}
}
