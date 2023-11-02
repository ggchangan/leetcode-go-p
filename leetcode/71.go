package leetcode

import (
	"strings"
)

var SimplifyPath = simplifyPath

func simplifyPath(path string) string {
	stack := make([]string, 0)
	pathes := strings.Split(path, "/")
	for i := range pathes {
		word := pathes[i]
		if len(word) == 0 || word == "." {
			continue
		} else if word == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, word)
		}
	}

	return "/" + strings.Join(stack, "/")
}
