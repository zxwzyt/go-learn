package ch8

import (
	"fmt"
	"os"
	"time"
)

func Countdown1() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("launch aborted!")
		return
	}

	// tick := time.Tick(1 * time.Second)
	// for countdown := 10; countdown > 0; countdown-- {
	// 	fmt.Println(countdown)
	// 	<-tick
	// }
	lanch()
}

func lanch() {
	fmt.Println("Lift off!")
}

func Select1() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
