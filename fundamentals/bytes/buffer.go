package bytes

import (
	"bytes"
	"fmt"
	"strings"
)

func RunInts() {
	fmt.Println(intsToStrings(1, 2, 3, 4, 5))
}
func intsToStrings(values ...int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// add comma each thre chars
func Commas(str string) string {
	number := strings.Split(str, ".")
	for i, substring := range number {

		var buf bytes.Buffer
		thisStr := substring
		for len(thisStr) > 0 {
			if len(thisStr) >= 4 {
				buf.WriteString(thisStr[:3] + ",")
				thisStr = thisStr[3:]
				continue
			}
			buf.WriteString(thisStr)
			thisStr = ""

		}
		number[i] = buf.String()
	}
	return strings.Join(number, ".")
}

func Anagrams(str1, str2 string) bool {
	isAnagram := true
	if len(str1) != len(str2) {
		return false
	}
	for _, val := range str1 {
		if !strings.Contains(str2, string(val)) {
			isAnagram = false
		}
	}
	return isAnagram
}
