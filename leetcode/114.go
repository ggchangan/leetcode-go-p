package leetcode

func flatten(root *TreeNode) {
	nodes := make([]*TreeNode, 0)
	var preOrder func(root *TreeNode)
	preOrder = func(root *TreeNode) {
		if root == nil {
			return
		}
		nodes = append(nodes, root)
		preOrder(root.Left)
		preOrder(root.Right)
	}

	preOrder(root)

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Left = nil
		nodes[i].Right = nodes[i+1]
	}
}
