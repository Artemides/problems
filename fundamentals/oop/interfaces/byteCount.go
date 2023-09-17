package interfaces

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func TestByteCounter() {
	var c ByteCounter
	c.Write([]byte("Hello"))
	fmt.Println(c)
	c = 0
	fmt.Fprintf(&c, "hello %s", "Sally")
	fmt.Println(c)

}

type MyReader interface {
	Read(doc string) (*html.Node, error)
}

type MyParser string

func (p *MyParser) Read(str string) (*html.Node, error) {
	reader := strings.NewReader(str)
	doc, err := html.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("parsing %q failed %s", str[:10], err)
	}

	return doc, nil
}

func RunReader() {
	var tmpl string
	var myParser MyParser
	tmpl = `
		<html>
		 <body>
		 <div>
		 	<p>A: </p>
			<p>B: </p>
		 </div>
		 </body>
		</html>
	`

	html, err := myParser.Read(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(html)
}
