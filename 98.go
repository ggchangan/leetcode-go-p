package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func helper(node *TreeNode, l, r int) bool {
	if node == nil {
		return true
	}

	if node.Val <= l || node.Val >= r {
		return false
	}

	return helper(node.Left, l, node.Val) && helper(node.Right, node.Val, r)
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}
