package concurrency

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Artemides/problems/fundamentals/anonymous"
	"golang.org/x/net/html"
)

func crawl(url string) []string {
	fmt.Printf("Crawling :\t%s\n", url)
	list, err := anonymous.Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}
func crawlV2(url string) []string {
	fmt.Printf("Crawling :\t%s\n", url)
	list, err := Extract(url)
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

var ctx, cancelRequest = context.WithCancel(context.Background())

func concurrentCrawlerWithCancel() {
	worklist := make(chan []string)
	unseen := make(chan string)
	seen := make(map[string]bool)

	depthLimit := flag.Uint64("depth", 20, "concurrent reachable url")
	flag.Parse()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("Crawler Cancelled")
		cancelRequest()
		close(worklist)
	}()

	go func() {
		urls := flag.Args()
		worklist <- urls
	}()

	for i := uint64(0); i < *depthLimit; i++ {
		go func() {
			for link := range unseen {
				go func(link string) {
					worklist <- crawlV2(link)
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

func Extract(url string) ([]string, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error Parsing %s as HTML %s", url, err)
	}

	var links []string
	visitNode := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key != "href" {
					continue
				}
				link, err := response.Request.URL.Parse(attr.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	anonymous.ForEachNode(doc, visitNode, nil)
	return links, nil
}

func CrawlerMain() {
	concurrentCrawlerWithCancel()
}
