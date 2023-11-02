package leetcode

var DiameterOfBinaryTree = diameterOfBinaryTree

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	maxHelper := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftDepth := depth(root.Left)
		rightDepth := depth(root.Right)

		diameter := leftDepth + rightDepth - 1
		if diameter > ans {
			ans = diameter
		}

		return maxHelper(leftDepth, rightDepth) + 1

	}

	depth(root)

	return
}
