package interfaces

import (
	"fmt"
	"io"
	"os"
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
