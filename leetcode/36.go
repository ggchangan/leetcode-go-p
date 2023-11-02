package leetcode

var IsValidSudoku = isValidSudoku

func isValidSudoku(board [][]byte) bool {
	checker := func(i, j int, sign []bool) bool {
		c := board[i][j]
		if c == '.' {
			return false
		}
		if sign[c-'0'] {
			return true
		}

		sign[c-'0'] = true
		return false
	}

	for i := 0; i < 9; i++ {
		sign := make([]bool, 10)
		for j := 0; j < 9; j++ {
			if checker(i, j, sign) {
				return false
			}
		}
	}

	for j := 0; j < 9; j++ {
		sign := make([]bool, 10)
		for i := 0; i < 9; i++ {
			if checker(i, j, sign) {
				return false
			}
		}
	}

	index := [3][2]int{
		{0, 3},
		{3, 6},
		{6, 9},
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sign := make([]bool, 10)
			for r := index[i][0]; r < index[i][1]; r++ {
				for c := index[j][0]; c < index[j][1]; c++ {
					if checker(r, c, sign) {
						return false
					}
				}
			}
		}
	}
	return true
}
