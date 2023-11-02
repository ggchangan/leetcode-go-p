package leetcode

var BuildTree2 = buildTree2

func buildTree2(inorder []int, postorder []int) *TreeNode {
	l := len(inorder)

	if l == 0 {
		return nil
	}
	if l == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}
	v := postorder[l-1]

	var i int
	for i = 0; i < l; i++ {
		if inorder[i] == v {
			break
		}
	}

	return &TreeNode{
		Val:   v,
		Left:  buildTree2(inorder[0:i], postorder[0:i]),
		Right: buildTree2(inorder[i+1:], postorder[i:l-1]),
	}
}
