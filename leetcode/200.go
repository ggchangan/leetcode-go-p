package leetcode

import "fmt"

func numIslands(grid [][]byte) int {
	var dfs func(i, j, m, n int)
	dfs = func(i, j, m, n int) {
		grid[i][j] = '0'
		fmt.Printf("%d,%d\n", i, j)

		if i-1 >= 0 {
			if grid[i-1][j] == '1' {
				dfs(i-1, j, m, n)
			}
		}
		if i+1 < m {
			if grid[i+1][j] == '1' {
				dfs(i+1, j, m, n)
			}
		}
		if j-1 >= 0 {
			if grid[i][j-1] == '1' {
				dfs(i, j-1, m, n)
			}
		}
		if j+1 < n {
			if grid[i][j+1] == '1' {
				dfs(i, j+1, m, n)
			}
		}
	}

	m := len(grid)
	n := len(grid[0])
	numbers := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(i, j, m, n)
				numbers += 1
			}
		}
	}

	return numbers
}
