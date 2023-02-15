package main

import (
	"fmt"
	"log"
	"os"

	"gobook/ch5/links"
)

// tokens is a semaphor counter used to impose the limit of 20 concurrent calls
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)
	var n int // number of pending send to the worklist

	n++
	// Start go routine to populate the worklist channel with the url's
	go func() { worklist <- os.Args[1:] }()

	// Start crawling concurrent
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
