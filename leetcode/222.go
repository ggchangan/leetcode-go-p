package leetcode

// methods 1. 2分查找
// 2. 递归满树
// 思路: 基本树节点个数 helper(root->left)+helper(root->right)+1
// 对于完全二叉树,如果一个树为满树，则可以用满树的计算公式
func countNodes(root *TreeNode) int {
	var helper func(root *TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := root.Left
		right := root.Right

		leftDepth, rightDepth := 0, 0
		for left != nil {
			leftDepth++
			left = left.Left
		}
		for right != nil {
			rightDepth++
			right = right.Right
		}

		if leftDepth == rightDepth {
			return (1 << (leftDepth + 1)) - 1
		}

		return helper(root.Left) + helper(root.Right) + 1
	}

	return helper(root)
}
