package dynamic

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		min := dp[i-1]
		for j := 2; j*j <= i; j++ {
			if dp[i-j*j] < min {
				min = dp[i-j*j]
			}
		}
		dp[i] = 1 + min
	}

	return dp[n]
}
