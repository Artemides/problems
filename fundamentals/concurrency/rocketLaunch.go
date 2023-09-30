package concurrency

import (
	"fmt"
	"os"
	"time"
)

func RocketLaunchMain() {
	RocketLaunch()
}

func RocketLaunch() {
	fmt.Println("Countdown to Launch, press returns to abort ")
	ticker := time.NewTicker(1 * time.Second)
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	defer ticker.Stop()

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-ticker.C:
		case <-abort:
			fmt.Println("Launch aborted")
			return
		}
	}

	<-time.After(10)
	fmt.Println("Launching ")
}
func RocketLaunchV2() {
	abort := make(chan struct{})
	fmt.Println("Countdown to Launch, press returns to abort ")

	start := time.Now()
	select {
	case <-time.After(3 * time.Second):
		fmt.Printf("cd: %0.fs\n", time.Since(start).Seconds())
	case <-abort:
		fmt.Println("aborting")
		return
	}
}
