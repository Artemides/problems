package initializer

import (
	"fmt"
	"time"
)

var pc [256]byte

func Run() {
	start1 := time.Now()
	fmt.Println("1: popCount:", PopCount(129))
	dur1 := time.Since(start1)
	start2 := time.Now()
	fmt.Println("2: popCount Iter:", PopCountIter(129))
	dur2 := time.Since(start2)

	fmt.Println("dur1: ", dur1, "dur2: ", dur2)

}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

}

func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}
func PopCountIter(x uint64) int {
	var count byte
	for i := 0; i < 8; i++ {

		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}
