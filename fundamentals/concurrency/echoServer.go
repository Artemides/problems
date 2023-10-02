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
	"sync"
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
		go echo(cnn, input.Text(), 1*time.Second)
	}

	cnn.Close()
}
func handleEchoConnV2(cnn net.Conn) {
	input := bufio.NewScanner(cnn)
	readlines := make(chan string)
	var wg sync.WaitGroup

	go func() {
		defer close(readlines)
		for input.Scan() {
			readlines <- input.Text()
		}
	}()

	for line := range readlines {
		wg.Add(1)
		go func(ln string) {
			defer wg.Done()
			echo(cnn, ln, 1*time.Second)
		}(line)
	}
	fmt.Println("waiting")
	wg.Wait()
	fmt.Println("continue")

	if tcpCnn, ok := cnn.(*net.TCPConn); ok {
		fmt.Println("Close Write")
		tcpCnn.CloseWrite()
		return
	}

	fmt.Println("Close All")
	cnn.Close()
}
func handleEchoConnV3(cnn net.Conn) {
	inputs := make(chan string)
	timeout := time.NewTimer(10 * time.Second)
	go func() {
		input := bufio.NewScanner(cnn)
		for input.Scan() {
			msg := input.Text()
			inputs <- msg
			timeout.Reset(10 * time.Second)
		}
		close(inputs)
	}()

	for {
		select {
		case msg := <-inputs:
			fmt.Println("received: ", msg)
			go echo(cnn, msg, 1*time.Second)
		case <-timeout.C:
			log.Printf("Connextion Closed: %s", cnn.LocalAddr().Network())
			cnn.Close()
			return
		}
	}

}

func EchoServerMain() {
	EchoServerV3()
}
func EchoServer() {
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

// EchoServerV2 Counts the number of clientes connected to it
func EchoServerV2() {
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

		go handleEchoConnV2(cnn)

	}
}

// this echo server disconects a client if no activity is
// received within t seconds
func EchoServerV3() {
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

		go handleEchoConnV3(cnn)

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
