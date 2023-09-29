package concurrency

import (
	"flag"
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

	depthLimit := flag.Uint64("depth", 20, "concurrent reachable url")
	flag.Parse()
	setFlag := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "depth" {
			setFlag = true
		}
	})

	go func() {
		fmt.Println("parsed: ", flag.Parsed())
		if setFlag {
			worklist <- os.Args[2:]
			return
		}
		worklist <- os.Args[1:]
	}()

	for i := uint64(0); i < *depthLimit; i++ {
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
