package leetcode

var FindSubstring = findSubstring

func findSubstring(s string, words []string) (ans []int) {
	wordL, wordCount, sL := len(words[0]), len(words), len(s)
	if sL < wordCount*wordL {
		return
	}

	total := make(map[string]int, wordCount)
	for _, w := range words {
		total[w]++
	}

	splitS := func(start int) (ans []string) {
		for i := start; i <= sL-wordL; i += wordL {
			word := s[i : i+wordL]
			ans = append(ans, word)
		}

		return
	}

	solve := func(sIndex int, splitWords []string) {
		start, end, splitL := 0, 0, len(splitWords)
		window := make(map[string]int, 0) //窗口是个map
		cnt := 0                          //有效单词数量
		//窗口是个map，窗口操作,
		for end < splitL {
			w := splitWords[end]
			if end-start+1 > wordCount { //左滑窗口
				sWord := splitWords[start]
				window[sWord]--
				start++
				if window[sWord] < total[sWord] {
					cnt--
				}
			}

			window[w]++
			if window[w] <= total[w] {
				cnt++
			}

			if cnt == wordCount {
				ans = append(ans, sIndex+start*wordL) //出错 start --> start *wordL
			}

			end++
		}
	}

	for i := 0; i < wordL; i++ {
		solve(i, splitS(i))
	}

	return
}
