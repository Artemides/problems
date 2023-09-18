package interfaces

import (
	"flag"
	"fmt"
	"time"
)

func Flags() {
	var period = flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("Sleeping for %v", *period)
	time.Sleep(*period)
	fmt.Println("Woke")

}
