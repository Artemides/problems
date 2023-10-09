package testing

import (
	"log"
	"os"
	"testing"

	"github.com/Artemides/problems/fundamentals/maps"
)

type testCase struct {
	word     string
	expected map[rune]int
}

var testCases = []testCase{
	{"aba", map[rune]int{'a': 2, 'b': 1}},
	{"a a", map[rune]int{'a': 2, ' ': 1}},
	{"a ", map[rune]int{'a': 1, ' ': 1}},
	{"aa", map[rune]int{'a': 2}},
	{"a δδ ε", map[rune]int{'a': 1, 'δ': 2, 'ε': 1, ' ': 2}},
	{"a;iio*&", map[rune]int{'a': 1, ';': 1, 'i': 2, 'o': 1, '*': 1, '&': 1}},
}

func TestCharcount(t *testing.T) {

	defer func(oldStdin *os.File) {
		os.Stdin = oldStdin
	}(os.Stdin)

	for _, testCase := range testCases {
		r, w, err := os.Pipe()
		if err != nil {
			log.Fatal(err)
		}

		os.Stdin = r
		_, err = w.Write([]byte(testCase.word))
		w.Close()
		if err != nil {
			log.Printf("wrinting: %s err: %s", testCase.word, err)
		}

		runeCount, _, _ := maps.RuneCount()
		for _rune, _count := range runeCount {
			if want := testCase.expected[_rune]; want != _count {
				t.Errorf(`case:%s => rune:%q , found: %d , expected: %d`, testCase.word, _rune, _count, want)
			}
		}
	}
}
