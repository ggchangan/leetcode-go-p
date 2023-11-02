package leetcode

import "strings"

var WordPattern = wordPattern

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}
	pToS := map[byte]string{}
	sToP := map[string]byte{}

	for i := range words {
		if pToS[pattern[i]] != "" && pToS[pattern[i]] != words[i] || sToP[words[i]] > 0 && sToP[words[i]] != pattern[i] {
			return false
		}
		pToS[pattern[i]] = words[i]
		sToP[words[i]] = pattern[i]
	}

	return true
}
