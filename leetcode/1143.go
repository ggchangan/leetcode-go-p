package leetcode

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text2)
	n := len(text1)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	if text1[0] == text2[0] {
		dp[0][0] = 1
	}
	for j := 1; j < n; j++ {
		if text1[j] == text2[0] {
			dp[0][j] = 1
		} else {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < m; i++ {
		if text2[i] == text1[0] {
			dp[i][0] = 1
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[j] == text2[i] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[m-1][n-1]
}
