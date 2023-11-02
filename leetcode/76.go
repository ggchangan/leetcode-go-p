package leetcode

var MinWindow = minWindow

func minWindow(str string, t string) (ans string) {
	sL, tL := len(str), len(t)
	if sL < tL {
		return
	}

	data := map[byte]int{}
	for _, c := range []byte(t) {
		data[c]++
	}

	s, e := 0, 0
	window := map[byte]int{}
	cnt := 0
	minL := 0
	for e < sL {
		c := str[e]
		window[c]++
		if window[c] <= data[c] {
			cnt++
		}
		for cnt >= tL {
			curL := e - s + 1
			if minL == 0 || curL < minL {
				minL = curL
				ans = str[s : e+1]
			}
			c := str[s]
			window[c]--
			if window[c] < data[c] {
				cnt--
			}
			s++
		}
		e++
	}

	return
}
