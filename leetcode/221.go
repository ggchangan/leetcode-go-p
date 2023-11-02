package leetcode

var MaximalSquare = maximalSquare

func maximalSquare(matrix [][]byte) int {
	minHelper := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				dp[i][j] = 0
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
				if dp[i][j] > ans {
					ans = dp[i][j]
				}
				continue
			}
			dp[i][j] = minHelper(minHelper(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			if dp[i][j] > ans {
				ans = dp[i][j]
			}
		}
	}

	return ans * ans
}
