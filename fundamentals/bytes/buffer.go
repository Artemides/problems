package bytes

import (
	"bytes"
	"fmt"
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
