package leetcode

import (
	"sort"
)

var GroupAnagrams = groupAnagrams

func groupAnagrams(strs []string) (ans [][]string) {
	m := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		str := []byte(strs[i])
		sort.Slice(str, func(i, j int) bool {
			return str[j] < str[i]
		})
		sortedStr := string(str)
		if _, ok := m[sortedStr]; !ok {
			m[sortedStr] = []string{strs[i]}
		} else {
			m[sortedStr] = append(m[sortedStr], strs[i])
		}
	}

	for _, v := range m {
		ans = append(ans, v)
	}

	return
}
