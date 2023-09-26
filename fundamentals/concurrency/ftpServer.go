package concurrency

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
)

func FTPServerMain() {
	serve()
}

func execCmd(w io.Writer, e string, args ...string) {
	cmd := exec.Command(e, args...)
	cmd.Stdout = w
	if err := cmd.Run(); err != nil {
		log.Print(err)
	}
}
func handleConnection(cnn net.Conn) {
	defer cnn.Close()

	input := bufio.NewScanner(cnn)
	for input.Scan() {
		cmds := strings.Split(input.Text(), " ")

		switch cmds[0] {
		case "ls":
			execCmd(cnn, cmds[0], cmds[1:]...)
		case "cd":
			if err := os.Chdir(cmds[1]); err != nil {
				log.Print(err)
			}
		case "get":
			file, err := os.Open(cmds[1])
			if err != nil {
				log.Printf("by opening %s  error: %v", cmds[1], err)
				continue
			}
			handleWrite(cnn, file)
		case "close":
			return
		default:
			help := "-ls listContent\n-cd change dir\n-get get content file\n-close close conection\n"
			handleWrite(cnn, strings.NewReader(help))
		}
	}
}

func serve() {
	port := flag.String("port", "3000", "port")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		cnn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error connectin %s", cnn)
			continue
		}
		go handleConnection(cnn)
	}
}
