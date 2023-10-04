package concurrency

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	msgChan chan<- string
	name    string
}

var (
	incomingClients = make(chan client)
	leavingClients  = make(chan client)
	messages        = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			for client := range clients {
				client.msgChan <- msg
			}
		case client := <-incomingClients:
			clients[client] = true
			go printClientsSet(messages, clients, client)

		case client := <-leavingClients:

			delete(clients, client)
			close(client.msgChan)
		}

	}
}

func printClientsSet(messages chan<- string, clients map[client]bool, joinedClient client) {
	var format string
	var idx int = 0
	for client := range clients {
		idx++
		if joinedClient.name == client.name {
			format += fmt.Sprintf("%d. %s - recently joined\n", idx, client.name)
			continue
		}

		format += fmt.Sprintf("%d. %s\n", idx, client.name)
	}
	messages <- format
}

func serveChat() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()

	for {
		cnn, err := listener.Accept()
		if err != nil {
			log.Printf("connection: %s, closed or failed: %s", cnn.RemoteAddr().String(), err)
			continue
		}
		go handleChatConnection(cnn)
	}
}

func handleChatConnection(cnn net.Conn) {
	inputsChan := make(chan string)
	timeout := time.NewTimer(15 * time.Second)
	ch := make(chan string)

	go handleWritingMessages(cnn, ch)

	me := cnn.RemoteAddr().String()
	ch <- "connected as : " + me
	incomingClients <- client{ch, me}

	go scanClientMessage(cnn, inputsChan, timeout)

loop:
	for {
		select {
		case msg := <-inputsChan:
			messages <- msg
		case <-timeout.C:
			fmt.Fprintf(cnn, "You were disconnected due to inactivity.")
			cnn.Close()
			break loop
		}
	}

	leavingClients <- client{ch, me}
	messages <- me + " : has left"

}

func scanClientMessage(cnn net.Conn, inputsChan chan<- string, timer *time.Timer) {
	input := bufio.NewScanner(cnn)
	me := cnn.RemoteAddr().String()

	for input.Scan() {
		inputsChan <- me + " : " + input.Text()
		timer.Reset(15 * time.Second)
	}
	close(inputsChan)
}

func handleWritingMessages(cnn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(cnn, msg)
	}
}
func ChatMain() {
	serveChat()
}
