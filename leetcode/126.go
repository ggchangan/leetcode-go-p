package leetcode

// TODO 不知道怎么记录广度优先搜索中产生的结果
func findLadders(beginWord string, endWord string, wordList []string) (ans [][]string) {
	wordAdj := make(map[string][]string, 0)
	valid := func(word1, word2 string) bool {
		count := 0
		for i := 0; i < len(word1); i++ {
			if word1[i] != word2[i] {
				count++
				if count > 1 {
					break
				}
			}
		}
		return count == 1
	}

	adj := func() {
		for _, word1 := range append(wordList, beginWord) {
			adjWords := make([]string, 0)
			for _, word2 := range wordList {
				if valid(word1, word2) {
					adjWords = append(adjWords, word2)
				}
			}
			wordAdj[word1] = adjWords
		}
	}
	adj()

	if _, ok := wordAdj[endWord]; !ok {
		return
	}
	if len(wordAdj[beginWord]) == 0 {
		return
	}

	cur := []string{beginWord}
	visited := make(map[string]bool, 0)
	for len(cur) > 0 {
		next := make([]string, 0)
		for i := 0; i < len(cur); i++ {
			if _, ok := visited[cur[i]]; ok {
				continue
			}
			if cur[i] == endWord {
				// TODO
				return
			}
			next = append(next, wordAdj[cur[i]]...)
			visited[cur[i]] = true
		}
		cur = next
	}

	return
}
