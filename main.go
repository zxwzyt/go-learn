package main

import "go-learn/training"

// func main() {
// 	images := []string{
// 		"/Users/anita/Desktop/a.jpg",
// 		"/Users/anita/Desktop/b.jpg",
// 		"/Users/anita/Desktop/c.jpg",
// 	}
// 	training.MakeThumbnails(images)
// }

// func main() {
// 	worklist := make(chan []string)

// 	go func() {
// 		worklist <- os.Args[1:]
// 	}()

// 	seen := make(map[string]bool)
// 	for list := range worklist {
// 		fmt.Println(111111)
// 		for _, link := range list {
// 			fmt.Println(22222)
// 			if !seen[link] {
// 				seen[link] = true
// 				go func(link string) {
// 					fmt.Println(333333)
// 					worklist <- training.Crawl(link)
// 				}(link)
// 			}
// 		}
// 	}
// }

// func main() {
// 	worklist := make(chan []string)
// 	var n int
// 	n++

// 	go func() {
// 		worklist <- os.Args[1:]
// 	}()

// 	seen := make(map[string]bool)

// 	for ; n > 0; n-- {
// 		list := <-worklist
// 		for _, link := range list {
// 			if !seen[link] {
// 				seen[link] = true
// 				n++
// 				go func(link string) {
// 					worklist <- training.Crawl1(link)
// 				}(link)
// 			}
// 		}
// 	}
// }

// func main() {
// 	worklist := make(chan []string)

// 	unseenLinks := make(chan string)

// 	go func() {
// 		worklist <- os.Args[1:]
// 	}()

// 	for i := 0; i < 20; i++ {
// 		go func() {
// 			for link := range unseenLinks {
// 				foundLinks := training.Crawl(link)
// 				go func() {
// 					worklist <- foundLinks
// 				}()
// 			}
// 		}()
// 	}

// 	seen := make(map[string]bool)
// 	for list := range worklist {
// 		for _, link := range list {
// 			if !seen[link] {
// 				seen[link] = true
// 				unseenLinks <- link
// 			}
// 		}
// 	}
// }

func main() {
	training.Countdown1()
}
