package deferred

import (
	"fmt"

	"github.com/Artemides/problems/fundamentals/anonymous"
	"golang.org/x/net/html"
)

func ReportMultipleHTMLTitleTags(document *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		//allows a program to handle the behavior of a panicking goroutine
		//if executed in a deferred function
		//	stops de panic sequence by restoring the normal execution and retrieving the error passed to the panic
		//if executed outside a defer
		// won't stop panic sequence and will return nil even tought the go routine doesn't panic
		switch p := recover(); p {
		case nil:
			//no panic
		case bailout{}:
			//expected pani
			err = fmt.Errorf("multiple title Elements")
		default:
			panic(p) //unknown or unexpected panic
		}
	}()
	anonymous.ForEachNode(document, func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "title" && node.FirstChild != nil {
			if title != "title" {
				panic(bailout{})
			}
			title = node.FirstChild.Data
		}
	}, nil)

	if title == "" {
		return title, fmt.Errorf("no title element found")
	}
	return title, nil
}
