package leetcode

var PathSum = pathSum

func pathSum(root *TreeNode, targetSum int) (ans int) {

	if root == nil {
		return
	}

	var rootSum func(root *TreeNode, targetSum int) int
	rootSum = func(root *TreeNode, targetSum int) (ans int) {
		if root == nil {
			return
		}

		if root.Val == targetSum {
			ans = 1
		}

		ans += rootSum(root.Left, targetSum-root.Val)
		ans += rootSum(root.Right, targetSum-root.Val)

		return
	}

	ans = rootSum(root, targetSum)

	ans += pathSum(root.Left, targetSum-root.Val)
	ans += pathSum(root.Right, targetSum-root.Val)

	return
}
