package leetcode

var LongestPalindrome = longestPalindrome

// 思路：定义函数p(i,j)代表s[i..j]是否为回文串，则p(i,j)=p(i-1,j+1)&&s[i]==s[j]
func longestPalindrome(s string) (ans string) {
	l := len(s)
	if l == 1 {
		return s
	}
	if l == 2 {
		if s[0] == s[1] {
			return s
		}
		return s[1:]
	}

	dp := make([][]bool, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]bool, l)
		dp[i][i] = true
	}
	ans = s[0:1]
	for i := 0; i < l-1; i++ {
		dp[i][i+1] = s[i] == s[i+1]
		if dp[i][i+1] {
			ans = s[i : i+2]
		}
	}

	for d := 3; d <= l; d++ { //对角线填充
		for start := 0; start < l-d+1; start++ {
			end := d + start - 1
			dp[start][end] = dp[start+1][end-1] && s[start] == s[end]
			if dp[start][end] {
				ans = s[start : end+1]
			}
		}
	}

	return
}
