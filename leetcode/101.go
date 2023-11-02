package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var helper func(p *TreeNode, q *TreeNode) bool

	helper = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if (p == nil && q != nil) || (p != nil && q == nil) {
			return false
		}

		return (p.Val == q.Val) && helper(p.Left, q.Right) && helper(p.Right, q.Left)
	}

	return helper(root.Left, root.Right)
}
