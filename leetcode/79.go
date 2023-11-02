package leetcode

var Exist = exist

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	l := len(word)
	var dfs func(i, j, p int) bool
	dfs = func(i, j, p int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == '#' || board[i][j] != word[p] {
			return false
		}

		if p == l-1 {
			return true
		}

		c := board[i][j]
		board[i][j] = '#'
		if dfs(i-1, j, p+1) || dfs(i+1, j, p+1) || dfs(i, j-1, p+1) || dfs(i, j+1, p+1) {
			return true
		}
		board[i][j] = c

		return false
	}

	//if dfs(1, 3, 0) {
	//	return true
	//}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != word[0] {
				continue
			}

			if dfs(i, j, 0) {
				return true
			}
		}
	}

	return false
}
