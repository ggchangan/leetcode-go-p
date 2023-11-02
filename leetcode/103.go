package leetcode

var ZigzagLevelOrder = zigzagLevelOrder

func zigzagLevelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}

	ans = make([][]int, 0)
	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		l := len(queue)
		//访问
		ans = append(ans, []int{})
		if level%2 == 0 {
			for i := 0; i < l; i++ {
				ans[level] = append(ans[level], queue[i].Val)
			}
		} else {
			for i := l - 1; i >= 0; i-- {
				ans[level] = append(ans[level], queue[i].Val)
			}
		}

		cur := make([]*TreeNode, 0)
		for i := 0; i < l; i++ {
			e := queue[i]
			if e.Left != nil {
				cur = append(cur, e.Left)
			}
			if e.Right != nil {
				cur = append(cur, e.Right)
			}
		}

		queue = cur
		level++
	}

	return
}
