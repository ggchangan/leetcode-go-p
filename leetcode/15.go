package leetcode

import (
	"sort"
)

var ThreeSum = threeSum

func threeSum(nums []int) (ans [][]int) {
	a := sort.IntSlice(nums[0:])
	sort.Sort(a)

	twoSum := func(nums []int, s, target int) {
		l, r := s+1, len(nums)-1
		for l < r {
			t := nums[l] + nums[r]
			if t == target {
				ans = append(ans, []int{nums[s], nums[l], nums[r]})
				old := l
				for l < r && nums[l] == nums[old] {
					l++
				}
			} else if t < target {
				l++
			} else {
				r--
			}
		}
	}

	for i, v := range a {
		if i > 0 && a[i] == a[i-1] {
			continue
		}
		twoSum(a, i, -1*v)
	}

	return ans
}
