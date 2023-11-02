package leetcode

var FindAnagrams = findAnagrams

func findAnagrams(str string, p string) (ans []int) {
	sL, pL := len(str), len(p)
	if sL < pL {
		return
	}

	m := map[byte]int{}
	for i := 0; i < pL; i++ {
		m[p[i]]++
	}

	s, e, validL := 0, 0, 0
	window := map[byte]int{}

	for e < sL {
		c := str[e]
		wL := e - s + 1
		if wL > pL {
			sC := str[s]
			window[sC]--
			if window[sC] < m[sC] {
				validL--
			}
			s++
		}

		window[c]++
		if window[c] <= m[c] {
			validL++
		}
		if validL == pL {
			ans = append(ans, s)
		}

		e++
	}

	return
}
