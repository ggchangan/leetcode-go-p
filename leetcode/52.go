package leetcode

var TotalNQueens = totalNQueens

func totalNQueens(n int) (ans int) {
	var dfs func(l int)
	col := make([]bool, n)
	pie := make([]bool, 2*n)
	na := make([]bool, 2*n)
	dfs = func(l int) {
		if l == n {
			ans++
			return
		}
		for j := 0; j < n; j++ {
			if col[j] || pie[l+j] || na[l-j+n] {
				continue
			}
			col[j] = true
			pie[l+j] = true
			na[l-j+n] = true

			dfs(l + 1)

			col[j] = false
			pie[l+j] = false
			na[l-j+n] = false
		}

	}

	dfs(0)

	return
}
