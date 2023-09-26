package concurrency

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

func serveTime(addrs string, column int, ch chan<- string) {
	cnn, err := net.Dial("tcp", addrs)
	if err != nil {
		log.Fatal(err)
	}
	defer cnn.Close()
	handleTime(cnn, column, ch)
}

func handleTime(in io.Reader, column int, ch chan<- string) {
	buf := make([]byte, 1024)
	for {
		n, err := in.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}
		data := string(buf[:n])

		ch <- fmt.Sprintf("%v&%s", column, data)
	}
}

func WriteTable(ch chan string, locations []string) {
	var format string
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	for range locations {
		format += "%s\t\n"
	}
	var locs []string
	for _, loc := range locations {
		locName := strings.Split(loc, "=")[0]
		locs = append(locs, locName)
	}

	fmt.Fprintf(tw, format, locs)
	tw.Flush()
	for data := range ch {
		//col&time
		metaData := strings.Split(data, "&")
		columns, _ := strconv.Atoi(metaData[0])
		time := metaData[1]
		spaces := make([]string, len(locations))
		spaces[columns] = time
		fmt.Fprintf(tw, format, spaces)
	}
}

func ClockWallMain() {
	servers := os.Args[1:]

	channel := make(chan string)

	go WriteTable(channel, servers)

	for column, server := range servers[1:] {
		addrs := strings.Split(server, "=")
		addr := addrs[1]
		go serveTime(addr, column+1, channel)
	}
	addr := strings.Split(servers[0], "=")[1]
	serveTime(addr, 0, channel)
	defer close(channel)
}
