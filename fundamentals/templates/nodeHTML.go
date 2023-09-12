package templates

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

type Node struct {
	Type                      NodeType
	Data                      string
	Attributes                []Attribute
	FirstChild, NextSibbbling *Node
}

type NodeType int32

type Attribute struct {
	Key, Val string
}

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DocTypeNode
)

// func ParseNode(r io.Reader) (*Node, error)

func RunNodeHTML() {

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parsing Err: %v", err)
		os.Exit(1)
	}
	links := visit(nil, doc)
	fmt.Printf("links: %d\n", len(links))
	for _, link := range links {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil || (n.Parent == nil && n.FirstChild == nil) {
		return links
	}

	if n.FirstChild == nil {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}

		parent := n.Parent
		parent.RemoveChild(n)
		return visit(links, parent)
	}
	return visit(links, n.FirstChild)
}

// func visit(links []string, n *html.Node) []string {
// 	if n.Type == html.ElementNode && n.Data == "a" {
// 		for _, a := range n.Attr {
// 			if a.Key == "href" {
// 				links = append(links, a.Val)
// 			}
// 		}
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		links = visit(links, c)
// 	}
// 	return links
// }

func PrintHtmlOutline() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parsing err: %s", err)
	}
	outline(nil, doc)

}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {

		outline(stack, c)
	}
}
