package recursion

func Sum(nums []int) []int {
	if len(nums) == 0 || len(nums) == 1 {
		return nums
	}
	res := nums[:len(nums)-1]
	nums[len(nums)-1] += Sum(res)[len(res)-1]
	return nums
}
