package leetcode

func sumNumbers(root *TreeNode) int {
	numbers := make([]int, 0)
	var preOrder func(root *TreeNode, sum int)
	preOrder = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}

		sum = sum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			numbers = append(numbers, sum)
			return
		}

		preOrder(root.Left, sum)
		preOrder(root.Right, sum)

	}
	preOrder(root, 0)

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}
