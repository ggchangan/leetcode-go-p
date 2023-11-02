package leetcode

import "sort"

var MergeIntervals = mergeIntervals

func mergeIntervals(intervals [][]int) (ans [][]int) {
	maxHelper := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	if len(intervals) == 1 {
		ans = intervals
		return
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		l := len(ans)
		_, rCur := ans[l-1][0], ans[l-1][1]
		if intervals[i][0] > rCur {
			ans = append(ans, intervals[i])
		} else {
			ans[l-1][1] = maxHelper(rCur, intervals[i][1])
		}
	}

	return
}
