package leetcode

var SetZeroes = setZeroes

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	row0, col0 := false, false
	for j := 0; j < n; j++ { //row0
		if matrix[0][j] == 0 {
			row0 = true
			break
		}
	}

	for i := 0; i < m; i++ { //col0
		if matrix[i][0] == 0 {
			col0 = true
			break
		}
	}

	// check 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if row0 {
		for j := 0; j < n; j++ { //row0
			matrix[0][j] = 0
		}
	}

	if col0 {
		for i := 0; i < m; i++ { //col0
			matrix[i][0] = 0
		}
	}
}
