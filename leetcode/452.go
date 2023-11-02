package leetcode

import "sort"

var FindMinArrowShots = findMinArrowShots

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	minHelper := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	maxHelper := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	ans := [][]int{points[0]}
	for i := 1; i < len(points); i++ {
		l := len(ans)
		if points[i][0] > ans[l-1][1] {
			ans = append(ans, points[i])
		} else {
			ans[l-1] = []int{maxHelper(ans[l-1][0], points[i][0]), minHelper(ans[l-1][1], points[i][1])}
		}
	}

	return len(ans)
}
