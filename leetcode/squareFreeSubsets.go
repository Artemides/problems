package leetcode

import (
	"fmt"
	"math"
)

func SquareFreeSubsets(nums []int) int {
	combinations := [][]int{}
	for idx := 0; idx < len(nums); idx++ {
		num := nums[idx]
		if !isSquareFree(num) {
			nums = append(nums[:idx], nums[idx+1:]...)
			idx--
		}
	}
	fmt.Println("nums", (nums))

	combine(nums, 0, []int{}, &combinations)
	// combine([]int{3, 4, 4, 5}, 0, []int{}, &combinations)
	// fmt.Println("combinations", (combinations))
	fmt.Println("combinations", len(combinations))
	fmt.Println(len(combinations) % ((1000000000) + 7))
	return len(combinations) % ((1000000000) + 7)
}

func isSquareFree(n int) bool {
	if n <= 1 {
		return true
	}

	// Factorize n into its prime factors
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			// Check if i is a repeated factor
			if n%(i*i) == 0 {
				return false
			}

			// Continue dividing n by i until it's no longer divisible
			for n%i == 0 {
				n /= i
			}
		}
	}

	// If n is greater than 1 at this point, it's a prime factor
	if n > 1 {
		if n%(n*n) == 0 {
			return false
		}
	}

	return true
}
func combine(nums []int, index int, subset []int, combination *[][]int) {
	if index == len(nums) {
		if len(subset) > 0 {
			mul := mulSlice(subset)
			if isSquareFree(mul) {
				*combination = append(*combination, subset)
			}
		}

		return

	}

	combine(nums, index+1, append(subset, nums[index]), combination)
	combine(nums, index+1, subset, combination)

}

func mulSlice(slice []int) int {
	if len(slice) == 0 {
		return 0
	}

	total := 1
	for _, v := range slice {
		total *= v
	}
	return total
}
