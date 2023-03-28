package training

import (
	"fmt"
	"log"
	"zxw-go/ch5"
)

var tokens = make(chan struct{}, 20)

func Crawl1(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := ch5.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
