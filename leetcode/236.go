package leetcode

// 思路：
// 递归解法：flson代表根节点x的左子树包含p或者q，frson代表根节点x的右子树包含p或者q
// 则x为p或者q的最近公共祖先满足: (flson && frson) || ((x == p || x == q)&&(flson || frson))
func lowestCommonAncestor(root, p, q *TreeNode) (ans *TreeNode) {
	var helper func(root, p, q *TreeNode) bool
	helper = func(root, p, q *TreeNode) bool {
		if root == nil {
			return false
		}
		lson := helper(root.Left, p, q)
		rson := helper(root.Right, p, q)

		if (lson && rson) || ((root.Val == p.Val || root.Val == q.Val) && (lson || rson)) {
			ans = root
		}

		return lson || rson || (root.Val == p.Val) || (root.Val == q.Val)
	}

	helper(root, p, q)
	return
}

// 方法2，构造每个节点的father，然后，从下向上遍历
func lowestCommonAncestor2(root, p, q *TreeNode) (ans *TreeNode) {
	father := map[*TreeNode]*TreeNode{
		root: nil,
	}
	var helper func(root *TreeNode)
	helper = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			father[root.Left] = root
			helper(root.Left)
		}

		if root.Right != nil {
			father[root.Right] = root
		}

	}

	helper(root)

	fatherNodes := map[*TreeNode]bool{p: true}
	f := father[p]
	for f != nil {
		fatherNodes[f] = true
		f = father[f]
	}

	qF := q
	for qF != nil {
		if fatherNodes[qF] {
			ans = qF
			return
		}
		qF = father[qF]
	}

	return
}
