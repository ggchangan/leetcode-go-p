package leetcode

var AverageOfLevels = averageOfLevels

func averageOfLevels(root *TreeNode) (ans []float64) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		cur := make([]*TreeNode, 0)
		sum := 0.0
		for i := 0; i < l; i++ {
			e := queue[i]
			if e.Left != nil {
				cur = append(cur, e.Left)
			}
			if e.Right != nil {
				cur = append(cur, e.Right)
			}
			sum += float64(e.Val)
		}
		queue = cur
		ans = append(ans, sum/float64(l))
	}

	return
}
