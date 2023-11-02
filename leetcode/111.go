package leetcode

func helper2(l int, m int, node *TreeNode) int {
	if node == nil {
		return m
	}

	ml := helper2(l+1, m, node.Left)

	mr := helper2(l+1, m, node.Right)

	if ml == 0 {
		m = mr + 1
		return m
	}

	if mr == 0 {
		m = ml + 1
		return m
	}

	if ml < mr {
		m = ml + 1
	} else {
		m = mr + 1
	}

	return m
}
func minDepth(root *TreeNode) int {
	return helper2(0, 0, root)
}
