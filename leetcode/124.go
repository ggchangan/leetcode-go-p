package leetcode

import "math"

var MaxPathSum = maxPathSum

// 思路：定义f(x)为包含x的最大路径和，则递归定义函数
// 对于每个树，可以找到一个最大路径，同时要返回一个不同时包含左右子树的路径以与根组成路径
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt

	maxHelper := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return math.MinInt
		}

		lSum := helper(root.Left)
		rSum := helper(root.Right)

		//包含根节点的，此树中的最大路径
		tMax := root.Val
		if lSum > 0 {
			tMax += lSum
		}
		if rSum > 0 {
			tMax += rSum
		}
		if tMax > maxSum {
			maxSum = tMax
		}

		//不同时包含左右子树的最大路径
		return root.Val + maxHelper(0, maxHelper(lSum, rSum))
	}

	helper(root)

	return maxSum
}
