package variadicfuncs

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func fetchAndParseURL(url string) (*html.Node, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching %s: %s", url, err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching %s: %s", url, response.Status)
	}

	defer response.Body.Close()
	doc, err := html.Parse(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error Parsing as HTML %s: %s", url, err)
	}

	return doc, nil
}

func ElementsByTagName(tags ...string) {
	const url = "https://golang.org"
	if len(tags) <= 0 {
		log.Fatal("no HTML tags specified...")
		return
	}

	matches := make(map[string]int)

	doc, err := fetchAndParseURL(url)
	if err != nil {
		log.Fatal(err)
	}
	findTags(doc, matches, tags)
	for tag, elements := range matches {
		fmt.Printf("\t <%s> : %d\n", tag, elements)

	}
}
func includes(strings []string, target string) bool {
	if len(strings) <= 0 {
		return false
	}
	if strings[0] != target {
		return includes(strings[1:], target)
	}
	return true
}
func findTags(node *html.Node, matches map[string]int, tags []string) {

	if node.FirstChild == nil {
		if node.Type == html.ElementNode && includes(tags, node.Data) {
			matches[node.Data]++

		}
		parent := node.Parent
		if parent == nil {
			return
		}
		parent.RemoveChild(node)
		findTags(parent, matches, tags)
		return
	}

	findTags(node.FirstChild, matches, tags)

}
