package leetcode

import "fmt"

func RemoveDuplicatesII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	frequency := 0
	first := nums[0]
	for idx := 0; idx < len(nums); idx++ {
		num := nums[idx]
		if first == num {
			frequency++
		} else {
			first = num
			frequency = 1
		}

		if frequency > 2 {
			nums = append(nums[:idx-1], nums[idx:]...)
			idx--
		}
	}
	fmt.Println(nums)
	return len(nums)
}
