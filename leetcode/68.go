package leetcode

import "strings"

var FullJustify = fullJustify

func fullJustify(words []string, maxWidth int) (ans []string) {
	groupWords := func(width int) (groups [][]int) {
		l, r, sumWordLen := 0, 0, 0
		spaceLen := 0
		for r < len(words) {
			wordLen := len(words[r])
			if width >= wordLen+spaceLen {
				r++
				width -= wordLen + spaceLen
				sumWordLen += wordLen
				spaceLen = 1
			} else {
				width = maxWidth
				groups = append(groups, []int{l, r - 1, maxWidth - sumWordLen})
				l = r
				sumWordLen = 0
				spaceLen = 0
			}
		}

		if l != r { //r一直往前跑，最后r跑出界，此时还有一组应该加入结果集
			groups = append(groups, []int{l, r - 1, maxWidth - sumWordLen})
		}
		return
	}

	leftAlign := func(l, r, spaceCount int) (ans string) {
		ans = words[l]
		for i := 1; i <= r-l; i++ {
			ans += " " + words[i+l]
			spaceCount--
		}
		ans += strings.Repeat(" ", spaceCount)
		return
	}
	midAlign := func(l, r, spaceCount int) (ans string) {
		spaceGroupCount := r - l
		m := spaceCount % spaceGroupCount
		n := spaceCount / spaceGroupCount
		ans = words[l]
		for i := 1; i <= r-l; i++ {
			if i <= m {
				ans += strings.Repeat(" ", n+1) + words[l+i]
			} else {
				ans += strings.Repeat(" ", n) + words[l+i]
			}
		}
		return
	}

	groups := groupWords(maxWidth)
	l := len(groups)
	for i := 0; i < l-1; i++ {
		l, r, spaceCount := groups[i][0], groups[i][1], groups[i][2]
		if r == l {
			ans = append(ans, leftAlign(l, r, spaceCount))
		} else {
			ans = append(ans, midAlign(l, r, spaceCount))
		}
	}

	ans = append(ans, leftAlign(groups[l-1][0], groups[l-1][1], groups[l-1][2]))

	return
}
