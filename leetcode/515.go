package leetcode

import "math"

func largestValues(root *TreeNode) (values []int) {
	if root == nil {
		return
	}

	cur := []*TreeNode{root}
	for len(cur) > 0 {
		m := math.MinInt64
		next := make([]*TreeNode, 0)
		for i := 0; i < len(cur); i++ {
			if cur[i].Val > m {
				m = cur[i].Val
			}
			if cur[i].Left != nil {
				next = append(next, cur[i].Left)
			}
			if cur[i].Right != nil {
				next = append(next, cur[i].Right)
			}
		}
		values = append(values, m)
		cur = next
	}

	return
}
