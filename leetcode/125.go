package leetcode

import "strings"

var IsPalindrome = isPalindrome

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	helper := func(c byte) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
	}

	l, r := 0, len(s)-1

	for l < r {
		for l < r && !helper(s[l]) {
			l++
		}
		for l < r && !helper(s[r]) {
			r--
		}

		if l < r {
			if s[l] != s[r] {
				return false
			} else {
				l++
				r--
			}
		}

	}

	return true
}
