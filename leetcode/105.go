package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
	preLen := len(preorder)
	inLen := len(inorder)
	var helper func(preL, preR, inL, inR int) *TreeNode
	helper = func(preL, preR, inL, inR int) *TreeNode {
		if inR == inL {
			return nil
		}
		if inR-inL == 1 {
			return &TreeNode{
				Val: inorder[inL],
			}
		}

		v := preorder[preL]
		var i int
		for i = inL; i < inR; i++ {
			if v == inorder[i] {
				break
			}
		}

		return &TreeNode{
			Val:   v,
			Left:  helper(preL+1, preL+1+i-inL, inL, i),
			Right: helper(preL+1+i-inL, preR, i+1, inR),
		}
	}

	return helper(0, preLen, 0, inLen)
}
