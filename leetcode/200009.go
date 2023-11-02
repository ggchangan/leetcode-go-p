package leetcode

var MaxLen = maxLen

// dp[i][j] i..j是否为合法括号
// dp = U(dp[i][s]) s = [i,j], 划分的时候步长至少是2
// j-i 最大的值
// 从下向上
func maxLen(data string) (ans int) {
	l := len(data)
	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
		dp[i][i] = false
	}
	//边界条件
	for i := 0; i < l-1; i++ {
		j := i + 1
		if data[i:j+1] == "()" {
			dp[i][j] = true
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			validL := j - i + 1
			if validL%2 != 0 {
				continue
			}
			for s := i + 1; s < j; s += 2 {
				if dp[i][s] && dp[s+1][j] {
					dp[i][j] = true
					//validL := j - i + 1
					if validL > ans {
						ans = validL
					}
					break
				}
			}
		}
	}

	return
}
