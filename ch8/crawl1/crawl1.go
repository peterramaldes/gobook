// Don't run this, this make my pc lock, to many go routines i think so
package main

import (
	"fmt"
	"log"
	"os"

	"gobook/ch5/links"
)

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)

	// Start go routine to populate the worklist channel with the url's
	go func() { worklist <- os.Args[1:] }()

	// Start crawling concurrent
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
