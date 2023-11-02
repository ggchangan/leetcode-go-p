package leetcode

var GetMinimumDifference = getMinimumDifference

func getMinimumDifference(root *TreeNode) (ans int) {
	numbers := make([]int, 0)
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			inOrder(root.Left)
		}
		numbers = append(numbers, root.Val)
		if root.Right != nil {
			inOrder(root.Right)
		}
	}

	inOrder(root)

	ans = 200001
	for i := 1; i < len(numbers); i++ {
		if numbers[i]-numbers[i-1] < ans {
			ans = numbers[i] - numbers[i-1]
		}
	}

	return
}

func getMinimumDifference2(root *TreeNode) (ans int) {
	ans = 200001
	pre := -1 //使用pre记录之前遍历的节点
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			inOrder(root.Left)
		}
		if pre != -1 && root.Val-pre < ans {
			ans = root.Val - pre
		}

		pre = root.Val

		if root.Right != nil {
			inOrder(root.Right)
		}
	}
	inOrder(root)
	return
}
