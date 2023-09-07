package bytes

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Run() {
	str := "hello, Καλος"
	for i, r := range str {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	fmt.Println("runes", utf8.RuneCountInString(str))
}

func Basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	fields := strings.Fields("a bc def ghij")
	fmt.Println(fields)
	return s
}
