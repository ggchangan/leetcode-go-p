package leetcode

// dpMax[i] = max{dpMax[i-1]*nums[i],nums[i], dpMin[i-1]*nums[i]}
// dpMin[i] = min{dpMax[i-1]*nums[i],nums[i], dpMin[i-1]*nums[i]}
func maxProduct(nums []int) int {
	dpMax := make([]int, len(nums))
	dpMin := make([]int, len(nums))
	dpMax[0] = nums[0]
	dpMin[0] = nums[0]

	maxHelper := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	minHelper := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := 1; i < len(nums); i++ {
		dpMax[i] = maxHelper(nums[i], maxHelper(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]))
		dpMin[i] = minHelper(nums[i], minHelper(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]))
	}

	m := dpMax[0]
	for i := 1; i < len(dpMax); i++ {
		if dpMax[i] > m {
			m = dpMax[i]
		}
	}

	return m
}

func maxProduct2(nums []int) int {
	dpMax := nums[0]
	dpMin := nums[0]
	ans := nums[0]

	maxHelper := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	minHelper := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := 1; i < len(nums); i++ {
		mx := dpMax
		mi := dpMin
		dpMax = maxHelper(nums[i], maxHelper(mx*nums[i], mi*nums[i]))
		dpMin = minHelper(nums[i], minHelper(mx*nums[i], mi*nums[i]))
		ans = maxHelper(ans, dpMax)
	}

	return ans
}
