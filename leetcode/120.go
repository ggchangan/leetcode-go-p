package leetcode

import "math"

// dp[i][j] = min{dp[i-1][j-1], dp[i-1][j]} + a[i][j]
func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = math.MaxInt
	}
	dp[1] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := len(triangle[i]) - 1; j >= 0; j-- { //invert
			if dp[j+1] < dp[j] {
				dp[j+1] = dp[j+1] + triangle[i][j]
			} else {
				dp[j+1] = dp[j] + triangle[i][j]
			}
		}
	}

	min := math.MaxInt
	for i := 0; i < len(dp); i++ {
		if dp[i] < min {
			min = dp[i]
		}
	}

	return min
}
