package training

import (
	"fmt"
	"log"
	"zxw-go/ch5"
)

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := ch5.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
