package leetcode

import "math"

var FindPeakElement = findPeakElement

func findPeakElement(nums []int) int {
	get := func(i int) int {
		if i < 0 || i >= len(nums) {
			return math.MinInt
		}
		return nums[i]
	}

	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) >> 1
		if get(m) > get(m-1) && get(m) > get(m+1) {
			return m
		}

		if get(m) < get(m+1) {
			l = m + 1
		} else {
			r = m - 1
		}

	}

	return 0
}
