package leetcode

var SnakesAndLadders = snakesAndLadders

var Pos = func(order, n int) (i int, j int) {
	i = n - (order-1)/n - 1

	if n%2 == 0 {
		if i%2 == 1 {
			j = (order - 1) % n
		} else {
			j = n - (order-1)%n - 1
		}
	} else {
		if i%2 == 0 {
			j = (order - 1) % n
		} else {
			j = n - (order-1)%n - 1
		}
	}
	return
}

func snakesAndLadders(board [][]int) int {
	minHelper := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}

	n := len(board)
	order := 1
	maxOrder := n * n
	queue := []int{order}
	visited := make(map[int]bool)
	step := 1
	for len(queue) > 0 {
		cur := make([]int, 0)
		l := len(queue)
		for i := 0; i < l; i++ {
			order := queue[i]
			for newOrder := order + 1; newOrder <= minHelper(order+6, maxOrder); newOrder++ {
				//梯子or蛇
				x, y := Pos(newOrder, n)
				curOrder := newOrder
				if board[x][y] != -1 {
					curOrder = board[x][y]
				}
				if curOrder == maxOrder {
					return step
				}
				if !visited[curOrder] {
					visited[curOrder] = true
					cur = append(cur, curOrder)
				}
			}
		}
		queue = cur
		step++
	}

	return -1
}
