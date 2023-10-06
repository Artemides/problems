package threads

import (
	"fmt"
	"time"
)

const STAGES = 10000000

func VSMain() {
	var pipeline [STAGES]chan string

	for stageIdx := range pipeline {
		pipeline[stageIdx] = make(chan string)
	}

	origin := pipeline[0]
	sink := pipeline[STAGES-1]

	for i := 0; i < STAGES-1; i++ {
		go func(i int) {
			pipeline[i+1] <- <-pipeline[i]
		}(i)
	}
	now := time.Now()
	origin <- "hello"
	<-sink
	fmt.Println("time: ", time.Since(now))
}
