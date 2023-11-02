package leetcode

var KthSmallest = kthSmallest

func kthSmallest(root *TreeNode, k int) (ans int) {
	cnt := 0
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			inOrder(root.Left)
		}
		cnt++
		if cnt == k {
			ans = root.Val
			return
		}

		if root.Right != nil {
			inOrder(root.Right)
		}
	}
	inOrder(root)

	return
}
