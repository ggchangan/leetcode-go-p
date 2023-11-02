package leetcode

var maxHelper = func(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func robHelper(nums []int, lo, hi int) int {
	if hi-lo == 1 {
		return nums[lo]
	}

	first, second := nums[lo], maxHelper(nums[lo], nums[lo+1])

	for i := lo + 2; i < hi; i++ {
		first, second = second, maxHelper(second, first+nums[i])
	}

	return second
}
func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return maxHelper(robHelper(nums, 0, len(nums)-1), robHelper(nums, 1, len(nums)))
}
