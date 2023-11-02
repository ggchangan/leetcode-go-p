package leetcode

var LongestValidParentheses = longestValidParentheses

func longestValidParentheses(s string) (ans int) {
	l := len(s)
	dp := make([]int, l)
	for i := 1; i < l; i++ {
		if s[i] != ')' {
			continue
		}
		if s[i-1] == '(' { // ..()
			if i >= 2 {
				dp[i] = dp[i-2] + 2
			} else {
				dp[i] = 2
			}
		} else { // .. )) 找到一种递推方式
			x := i - dp[i-1] - 1 //注意下标范围
			if x >= 0 && s[i-dp[i-1]-1] == '(' {
				if x-1 >= 0 {
					dp[i] = dp[i-1] + 2 + dp[x-1]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}

		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return
}
