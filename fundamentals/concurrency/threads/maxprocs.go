package threads

import "fmt"

func MAXPROCSMain() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
