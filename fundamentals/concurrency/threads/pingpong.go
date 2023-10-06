package threads

import (
	"fmt"
	"time"
)

func PingPongMain() {
	ping, pong := make(chan string), make(chan string)
	done := make(chan struct{})
	var sendingsPerSecond int
	timer := time.NewTimer(1 * time.Second)

	go func() {
		for {
			select {
			case <-done:
				return
			case v := <-ping:
				sendingsPerSecond++
				pong <- v
			}
		}

	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case v := <-pong:
				sendingsPerSecond++
				ping <- v
			}
		}
	}()

	ping <- "msg"
	<-timer.C
	close(done)

	select {
	case <-ping:
	case <-pong:
	}

	timer.Stop()
	fmt.Println("communitactions :", sendingsPerSecond)

}
