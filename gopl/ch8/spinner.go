package ch8

import (
	"fmt"
	"time"
)

// ============== goroutine ==================
func mainS() {
	go spinner(100 * time.Microsecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\r", n, fibN)
	time.Sleep(time.Second * 3)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
