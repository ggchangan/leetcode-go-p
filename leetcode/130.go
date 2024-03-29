package leetcode

var Solve = solve

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' { //对于不需要遍历的点直接返回
			return
		}

		board[i][j] = 'A'

		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			dfs(0, i)
		}
		if board[m-1][i] == 'O' {
			dfs(m-1, i)
		}
	}

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
}
