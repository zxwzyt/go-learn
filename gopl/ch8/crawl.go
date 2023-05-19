package ch8

import (
	"fmt"
	"go-learn/gopl/ch5"
	"log"
)

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := ch5.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
