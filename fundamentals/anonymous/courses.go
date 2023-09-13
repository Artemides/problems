package anonymous

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var prereq = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// compute a valid sequence of computer science courses prerequisites
func SortCourses(graph map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(nodes []string)
	visitAll = func(nodes []string) {
		for _, node := range nodes {
			if !seen[node] {
				seen[node] = true
				visitAll(graph[node])
				order = append(order, node)
			}
		}
	}

	var keys []string
	for _, key := range graph {
		keys = append(keys, key...)
	}

	visitAll(keys)
	return order
}

func TestCicle() {
	fmt.Println(courseCicled(&prereq, "formal languages", "formal languages"))
}
func courseCicled(graph *map[string][]string, targetCourse, currentCourse string) bool {
	courses := *graph
	prerequisites := courses[currentCourse]
	is := false
	if prerequisites == nil {
		return is
	}
	for _, prerequisite := range prerequisites {
		fmt.Println(prerequisite)
		if targetCourse == prerequisite {
			return true
		}

		is = is || courseCicled(graph, targetCourse, prerequisite)
	}
	return is
}
func SortCoursesMaps(graph map[string][]string) map[string][]string {
	order := make(map[string][]string)
	seen := make(map[string]bool)
	var visitAll func(nodes []string)
	visitAll = func(nodes []string) {
		for _, node := range nodes {
			if !seen[node] {
				seen[node] = true
				visitAll(graph[node])
				order[node] = append(order[node], graph[node]...)
			}
		}
	}

	var keys []string
	idx := 0
	for _, key := range graph {
		if idx >= 4 {
			break
		}
		keys = append(keys, key...)
		idx++
	}

	visitAll(keys)
	return order
}
func RunAnonymous() {
	order := SortCoursesMaps(prereq)
	for key, courses := range order {
		fmt.Printf("%s\n", key)
		for idx, course := range courses {
			fmt.Printf("\t%d. %s\n", idx+1, course)
		}

	}
}

func RunCrawl() {
	breadthFirst(crawl, os.Args[1:])
}

func Extract(url string) ([]string, error) {
	response, err := http.Get(url)
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
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(html *html.Node, pre, post func(html *html.Node)) {
	if pre != nil {
		pre(html)
	}

	for c := html.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(html)
	}

}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)

	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}

}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
