package leetcode

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxHelper := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	first, second := nums[0], maxHelper(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		first, second = second, maxHelper(second, first+nums[i])
	}

	return second
}
