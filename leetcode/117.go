package leetcode

type NextNode struct {
	Val   int
	Left  *NextNode
	Right *NextNode
	Next  *NextNode
}

var Connect = connect

func connect(root *NextNode) *NextNode {
	if root == nil {
		return root
	}

	queue := []*NextNode{root}
	for len(queue) > 0 {
		l := len(queue)
		cur := make([]*NextNode, 0)
		for i := 0; i < l-1; i++ {
			q := queue[i]
			if q.Left != nil {
				cur = append(cur, q.Left)
			}
			if q.Right != nil {
				cur = append(cur, q.Right)
			}
			q.Next = queue[i+1]
		}
		if queue[l-1].Left != nil {
			cur = append(cur, queue[l-1].Left)
		}
		if queue[l-1].Right != nil {
			cur = append(cur, queue[l-1].Right)
		}

		queue = cur
	}

	return root
}
