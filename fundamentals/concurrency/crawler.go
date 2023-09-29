package concurrency

import (
	"fmt"
	"log"
	"os"

	"github.com/Artemides/problems/fundamentals/anonymous"
)

func crawl(url string) []string {
	fmt.Printf("Crawling :\t%s\n", url)
	list, err := anonymous.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}

func ConcurrentCrawler() {
	worklist := make(chan []string)
	seen := make(map[string]bool)

	go func() { worklist <- os.Args[1:] }()

	for list := range worklist {
		for _, link := range list {
			if seen[link] {
				continue
			}

			seen[link] = true
			go func(link string) {
				worklist <- crawl(link)
			}(link)
		}
	}
}
func ConcurrentCrawlerV2() {
	worklist := make(chan []string)
	seen := make(map[string]bool)
	tokens := make(chan struct{}, 20)
	n := 0
	n++
	go func() { worklist <- os.Args[1:] }()

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if seen[link] {
				continue
			}

			seen[link] = true
			n++
			go func(link string) {
				tokens <- struct{}{}
				worklist <- crawl(link)
				<-tokens
			}(link)
		}
	}

}
func concurrentCrawlerV3() {
	worklist := make(chan []string)
	unseen := make(chan string)
	seen := make(map[string]bool)

	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseen {
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}()
	}

	for list := range worklist {
		for _, link := range list {
			if seen[link] {
				continue
			}

			seen[link] = true
			unseen <- link
		}
	}

}
func CrawlerMain() {
	concurrentCrawlerV3()
}
