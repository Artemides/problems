package concurrency

import (
	"fmt"
	"time"
)

func PipelineMain() {
	naturalChan := make(chan int)
	squareChan := make(chan int)

	go counter(naturalChan)
	go squarer(squareChan, naturalChan)
	printer(squareChan)
}

func counter(out chan<- int) {
	for i := 1; i <= 10; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}

func printer(in <-chan int) {
	for sq := range in {
		fmt.Println(sq)
	}
}

func SquarterPipeline2() {
	naturalChan := make(chan int)
	squareChan := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			naturalChan <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(naturalChan)
	}()

	go func() {
		for n := range naturalChan {
			squareChan <- n * n
		}
		close(squareChan)
	}()

	for sq := range squareChan {
		fmt.Println(sq)
	}

}

func SquarterPipeline() {
	naturalChan := make(chan int)
	squareChan := make(chan int)
	go func() {
		for i := 0; ; i++ {
			naturalChan <- i
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			square := <-naturalChan
			squareChan <- square * square
		}
	}()

	for {
		fmt.Println(<-squareChan)
	}

}
