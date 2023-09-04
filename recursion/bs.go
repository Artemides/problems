package recursion

import (
	"fmt"
	"math"
)

func Bs(arr []int, target, idx int) int {
	if len(arr) == 0 {
		return -1
	}

	if target != arr[0] && len(arr) == 1 {
		return -1
	}

	if target == arr[0] {
		return idx
	}

	fmt.Println(arr)
	m := int(math.Floor(float64(len(arr)) / 2))
	fmt.Println(m)
	if target < arr[m] {
		return Bs(arr[:m], target, idx)
	}
	return Bs(arr[m:], target, idx+m)

}
func Search() {
	arr := []int{1, 2, 5, 6, 7, 9, 12}
	index := Bs(arr, 9, 0)
	fmt.Println("index : ", index)
}
