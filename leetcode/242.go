package leetcode

var IsAnagram = isAnagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[byte]int{}
	for i := range s {
		m[s[i]] += 1
	}

	for i := range t {
		if m[t[i]] == 0 {
			return false
		}
		m[t[i]] -= 1
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
