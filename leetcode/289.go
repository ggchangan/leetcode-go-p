package leetcode

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])

	lives := func(i, j int) int {
		indexes := [8][2]int{
			{i - 1, j - 1},
			{i - 1, j},
			{i - 1, j + 1},
			{i, j - 1},
			{i, j + 1},
			{i + 1, j - 1},
			{i + 1, j},
			{i + 1, j + 1},
		}
		count := 0
		for _, indexes := range indexes {
			x, y := indexes[0], indexes[1]
			if x < 0 || x >= m || y < 0 || y >= n {
				continue
			}
			if board[x][y] > 0 {
				count++
			}
		}
		if board[i][j] == 0 {
			count *= -1
		}

		if count == 0 {
			return board[i][j]
		}
		return count
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = lives(i, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x := board[i][j]
			if x < 0 {
				if x == -3 {
					board[i][j] = 1
				} else {
					board[i][j] = 0
				}
			} else {
				if x < 2 || x > 3 {
					board[i][j] = 0
				}
				if x == 2 || x == 3 {
					board[i][j] = 1
				}
			}
		}
	}
}
