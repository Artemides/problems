package concurrency

import (
	"io"
	"log"
	"net"
	"time"
)

func ClockServerMain() {
	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		cnn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		handleConn(cnn)
	}

}
func handleConn(cnn net.Conn) {
	defer cnn.Close()
	for {
		_, err := io.WriteString(cnn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}

}
