package leetcode

var IsSubsequence = isSubsequence

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(t) == 0 {
		return false
	}

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		for j < len(t) && s[i] != t[j] {
			j++
		}
		if j < len(t) {
			i++
			j++
		}
	}

	return i == len(s)
}
