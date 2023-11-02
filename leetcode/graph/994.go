package graph

var OrangesRotting = orangesRotting

// 图的广度优先搜索
// 1.此图有多个起点，当做广度优先的第一层
// 2.将节点加入队列前，需要将节点设置为已经访问过了
func orangesRotting(grid [][]int) (ans int) {
	m := len(grid)
	n := len(grid[0])

	helper := func(stack [][]int) {
		for len(stack) > 0 {
			l := len(stack)
			curLevel := make([][]int, 0)
			for i := 0; i < l; i++ {
				cur := stack[i]
				x, y, step := cur[0], cur[1], cur[2]
				if step > ans {
					ans = step
				}
				for _, pos := range [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}} {
					nX, nY := pos[0], pos[1]
					if nX < 0 || nX >= m || nY < 0 || nY >= n {
						continue
					}
					if grid[nX][nY] != 1 {
						continue
					}
					grid[nX][nY] = -1 //放入队列之前，修改状态，避免多次放入
					curLevel = append(curLevel, []int{nX, nY, step + 1})
				}
			}
			stack = curLevel
		}
	}

	var stack [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 { //多个初始节点
				grid[i][j] = -1
				stack = append(stack, []int{i, j, 0})
			}
		}
	}
	helper(stack)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return

}
