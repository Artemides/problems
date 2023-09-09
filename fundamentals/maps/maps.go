package maps

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var mapping map[string]int

func Run() {
	baseMap := map[string]int{
		"rafa":     1,
		"rachelle": 2,
	}
	// mp := make(map[string]int)
	// baseSlc := make([]string, 0, len(mp))
	baseMap["ss"] = 1
}

func stringKey(slice []string) string {
	return fmt.Sprintf("%q", slice)
}

func Increase(slice []string) {
	key := stringKey(slice)
	mapping[key]++
}

func Count(slice []string) int {
	key := stringKey(slice)
	return mapping[key]
}

func RuneCount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stdout, "charcount: %v", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for key, count := range counts {
		fmt.Printf("%q\t%d\n", key, count)
	}
	fmt.Printf("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {

			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
