package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}

	cur := []*TreeNode{root}

	for len(cur) > 0 {
		values := make([]int, len(cur))
		next := make([]*TreeNode, 0)
		for i := 0; i < len(cur); i++ {
			values[i] = cur[i].Val
			if cur[i].Left != nil {
				next = append(next, cur[i].Left)
			}
			if cur[i].Right != nil {
				next = append(next, cur[i].Right)
			}
		}
		ans = append(ans, values)
		cur = next
	}

	return
}
