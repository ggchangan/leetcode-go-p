package leetcode

var CanConstruct = canConstruct

func canConstruct(ransomNote string, magazine string) bool {
	m := [26]int{0}
	for _, c := range magazine {
		m[c-'a'] += 1
	}

	for _, c := range ransomNote {
		m[c-'a'] -= 1
		if m[c-'a'] < 0 {
			return false
		}
	}

	return true
}
