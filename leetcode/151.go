package leetcode

var ReverseWords = reverseWords

func reverseWords(s string) (ans string) {
	p, q := 0, 0
	l := len(s)
	for p < l && q < l {
		//space
		for p < l && s[p] == ' ' {
			p++
		}
		q = p

		//word
		for q < l && s[q] != ' ' {
			q++
		}

		if q != p {
			if len(ans) == 0 {
				ans = s[p:q]
			} else {
				ans = s[p:q] + " " + ans
			}
		}
		p = q
	}

	return
}
