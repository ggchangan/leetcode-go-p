package leetcode

var InorderTraversal = inorderTraversal

func inorderTraversal(root *TreeNode) (ans []int) {
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			helper(root.Left)
		}

		ans = append(ans, root.Val)

		if root.Right != nil {
			helper(root.Right)
		}

	}

	helper(root)

	return
}
