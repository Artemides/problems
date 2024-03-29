package fundamentals

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CountFileLines2() {
	counts := make(map[string]int)
	lineFile := make(map[string]string)

	for _, filename := range os.Args[1:] {
		fileExt := filepath.Ext(filename)
		name := strings.TrimSuffix(filepath.Base(filename), fileExt)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if !strings.Contains(lineFile[line], name) {
				lineFile[line] += name + " "
			}
		}
	}
	for line, times := range counts {
		if times > 1 {
			files := lineFile[line]
			fmt.Printf("%v\t%s\t%s\n", times, files, line)
		}
	}
}

func CountFileLines() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
		return
	}

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}

		countLines(f, counts)
		f.Close()
	}
	for line, times := range counts {
		if times > 1 {
			fmt.Printf("%d\t%s\n", times, line)
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {

		counts[input.Text()]++
	}
}
