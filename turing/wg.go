package turing

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const N = 10

func WgMain() {
	m := make(map[int]int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println((m))

}

func Arr() {
	s2 := []int{1, 2, 3}
	s := s2[:0]
	fmt.Println(len(s), s, cap(s))
}

func MainSquares() {
	c := make(chan int, 3)
	go squares(c)
	fmt.Println(runtime.NumGoroutine())
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	time.Sleep(300 * time.Millisecond)
	fmt.Println(runtime.NumGoroutine())

}

func squares(c chan int) {
	for i := 0; i < 4; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func Printing() {
	const (
		a = 8
		b
		c
	)
	fmt.Println(a, b, c)
}

func ThreadMain() {
	fmt.Println(1)
	go func() {
		fmt.Println(2)
	}()

	select {}
	fmt.Println(3)
}
