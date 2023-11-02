package leetcode

import "strings"

func minDistance(word1 string, word2 string) int {
	m := len(word2)
	n := len(word1)

	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	if word1[0] == word2[0] {
		dp[0][0] = 0
	} else {
		dp[0][0] = 1
	}
	for i := 1; i < m; i++ {
		if strings.Contains(word2[:i+1], string(word1[0])) {
			dp[i][0] = i
		} else {
			dp[i][0] = i + 1
		}
	}

	for i := 1; i < n; i++ {
		if strings.Contains(word1[:i+1], string(word2[0])) {
			dp[0][i] = i
		} else {
			dp[0][i] = i + 1
		}
	}

	minHelper := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word2[i] == word1[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + minHelper(minHelper(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}
