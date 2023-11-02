package leetcode

var RightSideView = rightSideView

func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		ans = append(ans, queue[l-1].Val)
		cur := make([]*TreeNode, 0)
		for i := 0; i < l; i++ {
			e := queue[i]
			if e.Left != nil {
				cur = append(cur, e.Left)
			}
			if e.Right != nil {
				cur = append(cur, e.Right)
			}
		}
		queue = cur
	}

	return
}
