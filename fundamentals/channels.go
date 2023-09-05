package fundamentals

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func DoWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func Channels() {
	pipe := make(chan int)
	go func() {
		wg := sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				pipe <- DoWork()
			}()
		}
		wg.Wait()
		close(pipe)
	}()

	for data := range pipe {
		fmt.Println(data)
	}
}
