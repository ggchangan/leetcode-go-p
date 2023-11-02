package leetcode

var LengthOfLongestSubstring = lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	start, end := -1, 0
	maxL := 0
	m := make(map[byte]int, 0)
	for start < len(s) {
		if start >= 0 && m[s[start]] != 0 {
			delete(m, s[start])
		}
		for end < len(s) && m[s[end]] == 0 { //end 开始出现重复字符时退出

			m[s[end]] = 1
			end++
		}
		l := end - start - 1 //start指向的字符是要被排除在外
		if l > maxL {
			maxL = l
		}
		start++
	}
	return maxL
}
