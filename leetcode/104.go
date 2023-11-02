package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var m int

func helper3(l int, node *TreeNode) int {
	if node == nil {
		return m
	}

	if l > m {
		m = l
	}

	if node.Left != nil {
		helper3(l+1, node.Left)
	}

	if node.Right != nil {
		helper3(l+1, node.Right)
	}

	return m
}
func maxDepth(root *TreeNode) int {
	return helper3(1, root)
}
