package errors

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func RunErrors() {
	const url = "https://golang.orgasdas"
	if err := WaitForServer(url); err != nil {
		// fmt.FPrintf(os.Stderr,"Site is down: %s",err)
		// os.Exit(1)
		log.Fatalf("Site is down %s", err)
	}
}

func WaitForServer(url string) error {
	timeout := 15 * time.Second
	deadline := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Get(url)
		if err == nil {
			return nil
		}
		log.SetPrefix("wait:")
		log.SetFlags(0)
		log.Printf("Server Not Responding (%s); Retrying", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s ", url, timeout)
}
