package sharing

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func incommingURLs() []string {

	return []string{"https://golang.org", "https://play.golang.org", "http://gopl.io", "https://golang.org"}
}

func TestMemo(t *testing.T) {
	m := New(httpGetBody)
	for _, url := range incommingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}

		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}
