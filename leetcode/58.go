package leetcode

var LengthOfLastWord = lengthOfLastWord

func lengthOfLastWord(s string) (ans int) {
	p, q := 0, 0
	l := len(s)
	for p < l {
		for p < l && s[p] == ' ' {
			p++
		}
		q = p
		for p < l && s[p] != ' ' {
			p++
		}
		if p-q != 0 {
			ans = p - q
		}
	}

	return
}
