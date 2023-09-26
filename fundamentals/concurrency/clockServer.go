package concurrency

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func ClockServerMain() {
	var port = flag.String("port", "8000", "port number")
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listenning on : 127.0.0.1:" + *port)
	for {
		cnn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(cnn)
	}

}

func handleConn(cnn net.Conn) {
	for {
		_, err := io.WriteString(cnn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(1 * time.Second)
	}

}
