package concurrency

import (
	"io"
	"log"
	"net"
	"os"
)

func WriteServerMain() {
	cnn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	defer cnn.Close()

	handleWrite(os.Stdout, cnn)
}

func handleWrite(out io.Writer, in io.Reader) {
	if _, err := io.Copy(out, in); err != nil {
		log.Fatal(err)
	}
}
