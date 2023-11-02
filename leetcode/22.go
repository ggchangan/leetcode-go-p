package leetcode

var results []string

func parenthesis(l int, r int, n int, s string) []string {
	if l == n && r == n {
		results = append(results, s)
		return results
	}

	if l < n {
		parenthesis(l+1, r, n, s+"(")
	}

	if r < l {
		parenthesis(l, r+1, n, s+")")
	}

	return results
}

func generateParenthesis(n int) []string {
	results = []string{}
	return parenthesis(0, 0, n, "")
}
