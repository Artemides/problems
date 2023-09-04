package leetcode

import (
	"math"
	"sort"
)

func RemoveDuplicates(nums []int) int {
	uniques := make(map[int]bool)

	for idx, v := range nums {
		if _, ok := uniques[v]; !ok {
			uniques[v] = true
			nums[idx] = int(math.Inf(-1))
		}
	}
	sort.Ints(nums)
	return len(uniques)
}
