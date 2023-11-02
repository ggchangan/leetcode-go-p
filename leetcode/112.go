package leetcode

func hasPathSum(root *TreeNode, targetSum int) bool {
	sum := 0
	var preOrder func(root *TreeNode) bool
	preOrder = func(root *TreeNode) bool {
		if root == nil {
			return false
		}

		sum += root.Val

		if root.Left == nil && root.Right == nil && sum == targetSum {
			return true
		}

		if preOrder(root.Left) || preOrder(root.Right) {
			return true
		}

		sum -= root.Val

		return false
	}

	return preOrder(root)
}
