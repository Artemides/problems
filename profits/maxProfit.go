package profits

import (
	"fmt"
	"sort"
)

func MaxProfit(prices []int) int {
	prices = append(prices, -1)
	profit := findNearests(prices)

	return profit
}

func findNearests(nums []int) int {
	var indexes []int
	from, to := 0, 0
	for idx, num := range nums[:len(nums)-1] {
		if nums[idx+1] < num {
			subProfit := nums[to] - nums[from]
			indexes = append(indexes, subProfit)
			from = idx + 1
			to = idx
		}
		to++
	}
	sort.Slice(indexes, func(i, j int) bool {
		return indexes[i] > indexes[j]
	})

	fmt.Println(indexes)
	if len(indexes) > 2 {
		indexes = indexes[:2]
	}

	return sum(indexes)
}

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
