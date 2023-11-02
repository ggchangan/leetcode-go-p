package leetcode

func ladderLength(beginWord string, endWord string, wordList []string) int {
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
		return 0
	}
	if len(wordAdj[beginWord]) == 0 {
		return 0
	}

	cur := []string{beginWord}
	changes := 1
	//visited := make(map[string]bool, 0)
	for len(cur) > 0 {
		next := make([]string, 0)
		for i := 0; i < len(cur); i++ {
			//if _, ok := visited[cur[i]]; ok {
			//	continue
			//}
			if cur[i] == endWord {
				return changes
			}
			next = append(next, wordAdj[cur[i]]...)
			//visited[cur[i]] = true
			delete(wordAdj, cur[i])
		}
		cur = next
		changes++
	}

	return 0
}
