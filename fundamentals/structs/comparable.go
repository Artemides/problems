package structs

import (
	"fmt"
	"time"
)

type Student struct {
	ID    string
	name  string
	email string
	birth time.Weekday
}

func Compare() {
	S1 := Student{"01", "name1", "email", time.Sunday}
	S2 := Student{"01", "name1", "email", time.Sunday}
	fmt.Println("S1 == S2 : ", S1 == S2)
}

func MapKey() {
	type Address struct {
		hostname string
		port     int
	}

	hits := make(map[Address]int)
	hits[Address{"localhost", 8800}]++
}
