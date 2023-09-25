package interfaces

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func RunAssertion() {
	var w io.Writer = os.Stdout
	rw := w.(io.ReadWriter)
	fmt.Printf("w: T% rw: %T\n", w, rw)

	_, err := os.Open("/no_such_file.os")
	if err != nil {
		fmt.Println(os.IsNotExist(err))
	}

}

func FormatOneValye(x interface{}) string {
	if err, ok := x.(error); ok {
		return err.Error()
	}
	if err, ok := x.(fmt.Stringer); ok {
		return err.String()
	}
	return ""
}

func SQLQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x)
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	case string:
		return SQLString(x)
	default:
		panic(fmt.Sprintf("unespecified type: %T as %v", x, x))
	}

}
func SQLString(x interface{}) string {
	if _, ok := x.(string); !ok {
		panic("not a string")
	}
	return fmt.Sprint(x)

}

func DecodeXML() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string

	for {
		tok, err := dec.Token()
		if err != io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
