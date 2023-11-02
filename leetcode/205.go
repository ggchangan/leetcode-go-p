package leetcode

var IsIsomorphic = isIsomorphic

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[byte]byte, 0)
	m2 := make(map[byte]byte, 0)
	for i := 0; i < len(t); i++ {
		if m1[s[i]] > 0 && m1[s[i]] != t[i] || m2[t[i]] > 0 && m2[t[i]] != s[i] {
			return false
		}

		m1[s[i]] = t[i]
		m2[t[i]] = s[i]
	}

	return true
}
