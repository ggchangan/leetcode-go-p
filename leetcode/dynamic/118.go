package dynamic

func generate(numRows int) (ans [][]int) {
	ans = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				ans[i][j] = 1
				continue
			}
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}

	return
}
