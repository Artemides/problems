package fundamentals

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func Fetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") || !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("status: %v\n", response.StatusCode)
		_, err = io.Copy(os.Stdout, response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
