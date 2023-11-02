package leetcode

import (
	"fmt"
)

func summaryRanges(nums []int) (ans []string) {
	l := len(nums)
	if l == 0 {
		return
	}
	if l == 1 {
		ans = append(ans, fmt.Sprintf("%d", nums[0]))
		return
	}
	p, q := 0, 1
	for q < l {
		for q < l && nums[q] == nums[q-1]+1 {
			q++
		}
		if q-p == 1 {
			ans = append(ans, fmt.Sprintf("%d", nums[p]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[p], nums[q-1]))
		}
		p = q
		q++
	}

	if p < l {
		ans = append(ans, fmt.Sprintf("%d", nums[p]))
	}

	return
}
