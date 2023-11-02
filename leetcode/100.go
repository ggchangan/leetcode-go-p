package leetcode

var IsSameTree = isSameTree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var helper func(p *TreeNode, q *TreeNode) bool
	helper = func(p *TreeNode, q *TreeNode) bool {
		if (p == nil && q != nil) || (p != nil && q == nil) {
			return false
		}

		if p == nil && q == nil {
			return true
		}

		return (p.Val == q.Val) && helper(p.Left, q.Left) && helper(p.Right, q.Right)
	}

	return helper(p, q)
}
