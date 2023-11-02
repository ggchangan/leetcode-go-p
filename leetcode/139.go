package leetcode

import "strings"

func wordBreak(s string, wordDict []string) bool {
	visited := make(map[string]bool)

	var helper func(s string) bool
	helper = func(s string) bool {
		if _, ok := visited[s]; ok {
			return false
		}

		if len(s) == 0 {
			return true
		}

		for _, word := range wordDict {
			if strings.HasPrefix(s, word) {
				if helper(strings.TrimPrefix(s, word)) {
					return true
				}
			}
		}

		visited[s] = false
		return false
	}

	helper2 := func(s string) bool {
		dp := make([]bool, len(s)+1)
		dp[0] = true
		for i := 0; i < len(s); i++ {
			for j := 0; j < len(wordDict); j++ {
				if strings.HasSuffix(s[:i+1], wordDict[j]) {
					if dp[i+1-len(wordDict[j])] {
						dp[i+1] = true
						continue
					}
				}
			}
		}

		return dp[len(dp)-1]
	}

	//return helper(s)
	return helper2(s)
}
