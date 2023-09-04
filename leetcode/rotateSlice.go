package leetcode

import "fmt"

func SS() {
	// slc1 := []int{1, 2, 3, 4, 5, 6, 7}
	slc1 := []int{7, 2}
	rotations := 5

	rotations = rotations % len(slc1)
	slc2 := make([]int, len(slc1))
	copy(slc2, slc1[len(slc1)-rotations:])
	fmt.Println(slc2)
	copy(slc1[rotations:], slc1[:len(slc1)-rotations])
	copy(slc1, slc2[:rotations])
	fmt.Println(slc1)
}
