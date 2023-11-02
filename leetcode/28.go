package leetcode

var StrStr = strStr

func strStr(haystack string, needle string) int {
	hl := len(haystack)
	nl := len(needle)
	if hl < nl {
		return -1
	}

	p, q := 0, 0
	for p < hl && q < nl {
		for p < hl && q < nl && haystack[p] == needle[q] {
			if q == nl-1 {
				return p - nl + 1
			}
			p++
			q++
		}

	}

	return -1
}
