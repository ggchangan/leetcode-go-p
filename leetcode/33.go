package leetcode

var Search = search

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}

	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) >> 1
		if target == nums[m] {
			return m
		}

		//<= 重要, 开始写的是<， 此处容易出错
		if nums[l] <= nums[m] {
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target > nums[m] && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}

	}

	return -1
}
