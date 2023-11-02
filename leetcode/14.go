package leetcode

var LongestCommonPrefix = longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return strs[0]
	}
	var i int
	for i = 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
