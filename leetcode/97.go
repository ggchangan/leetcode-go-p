package leetcode

var IsInterleave = isInterleave

// 思路：f(i,j)代表s1的前i个字符和s2的前j个字符可以构成s3前i+j个字符的交错子串，则
// f(i, j) == (f(i-1, j) && s1[i-1]==s3[i+j-1]) or (f(i, j-1) && s2[j-1]==s3[i+j-1])
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1, l2, l3 := len(s1), len(s2), len(s3)
	if l3 != l1+l2 {
		return false
	}

	dp := make([][]bool, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]bool, l2+1)
	}

	dp[0][0] = true

	for i := 0; i <= l1; i++ {
		for j := 0; j <= l2; j++ {
			//这种方法正确初始化第一行和第一列并且不需要额外程序
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[i+j-1])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}

		}
	}

	return dp[l1][l2]
}
