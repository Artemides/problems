package leetcode

func RemoveElement(nums []int, val int) int {
	replacements := filter(nums, val)
	// sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	return len(nums) - replacements
}

func filter(nums []int, val int) int {
	times := 0
	for idx, v := range nums {
		if v == val {
			nums[idx] = -1
			times++
		}
	}
	return times
}
