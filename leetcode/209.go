package leetcode

import "math"

var MinSubArrayLen = minSubArrayLen

func minSubArrayLen(target int, nums []int) int {
	minHelper := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	sum := 0
	minL := math.MaxInt
	for s, e := 0, 0; e < len(nums); e++ {
		sum += nums[e]
		for sum >= target {
			minL = minHelper(minL, e-s+1)
			sum -= nums[s]
			s++
		}
	}

	if minL == math.MaxInt {
		return 0
	}

	return minL
}
