package interfaces

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func Flags() {
	var period = flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("Sleeping for %v", *period)
	time.Sleep(*period)
	fmt.Println("Woke")

}

func RunInterfaces() {
	var w io.Writer
	fmt.Printf("w value %T\n", w)
	w = os.Stdout
	fmt.Printf("w value %T\n", w)
	w = new(bytes.Buffer)
	fmt.Printf("w value %T\n", w)
	w = nil
	fmt.Printf("w value %T\n", w)

}

func NilInterfaces() {
	const debug = false
	// var buf *bytes.Buffer
	var buf io.Writer

	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf)

	if debug {
		fmt.Println("Debug")
	}

}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
