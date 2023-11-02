package leetcode

import "strings"

//1. pattern = "abba", str="北京 杭州 杭州 北京" 返回 ture
//2. pattern = "aabb", str="北京 杭州 杭州 北京" 返回 false
//3. pattern = "baab", str="北京 杭州 杭州 北京" 返回 ture

func CheckStrPattern(pattern string, str string) (ans bool) {
	pL := len(pattern)
	sL := len(str)

	if pL == 0 {
		ans = true
		return
	}

	if sL == 0 {
		ans = false
		return
	}

	words := strings.Split(str, " ")
	wL := len(words)
	if pL != wL {
		ans = false
		return
	}

	pToW := map[byte]string{}
	wToP := map[string]byte{}
	for i := 0; i < wL; i++ {
		p := pattern[i]
		w := words[i]
		if (len(pToW[p]) > 0 && pToW[p] != w) || (wToP[w] > 0 && wToP[w] != p) {
			ans = false
			return
		}

		pToW[p] = w
		wToP[w] = p
	}

	ans = true
	return
}
