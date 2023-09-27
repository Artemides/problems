package concurrency

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"time"
)

func echo(cnn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(cnn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(cnn, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(cnn, "\t", strings.ToLower(shout))

}

func handleEchoConn(cnn net.Conn) {
	input := bufio.NewScanner(cnn)

	for input.Scan() {
		go echo(cnn, input.Text(), 2*time.Second)
	}
	cnn.Close()
}

func EchoServerMain() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		cnn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection with %s , err: %s", cnn.LocalAddr().Network(), err)
			continue
		}

		go handleEchoConn(cnn)

	}
}

func EchoClient() {
	cnn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("Connection with %s , err: %s", cnn.LocalAddr().Network(), err)
	}

	defer cnn.Close()
	go handleWrite(os.Stdout, cnn)
	handleWrite(cnn, os.Stdin)
}
func EchoClientWait() {
	cnn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("Connection with %s , err: %s", cnn.LocalAddr().Network(), err)
	}

	done := make(chan interface{})
	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, os.Interrupt)

	go func() {
		io.Copy(os.Stdout, cnn)
		fmt.Println("DONE")
		done <- struct{}{}
	}()

	go handleWrite(cnn, os.Stdin)

	<-signChan
	if tcpCnn, ok := cnn.(*net.TCPConn); ok {
		fmt.Println("CLOSE WRITE")
		tcpCnn.CloseWrite()
	} else {
		fmt.Println("CLOSE ALL")
		cnn.Close()
	}
	fmt.Println("closed")
	<-done //wait for the go routine to complete its job
}
