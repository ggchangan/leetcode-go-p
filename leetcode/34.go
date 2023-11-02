package leetcode

var SearchRange = searchRange

func searchRange(nums []int, target int) (ans []int) {
	ans = []int{-1, -1}
	if len(nums) == 0 {
		return
	}

	find := false
	l, r := 0, len(nums)-1
	m := 0
	for l <= r {
		m = (l + r) >> 1
		if nums[m] == target {
			find = true
			break
		} else if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if !find {
		return
	}

	p, q := m, m
	for p >= 0 && nums[p] == target {
		p--
	}
	if p < 0 || nums[p] != target {
		p++
	}
	for q <= len(nums)-1 && nums[q] == target {
		q++
	}
	if q > len(nums)-1 || nums[q] != target {
		q--
	}

	ans = []int{p, q}
	return
}
