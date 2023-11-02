package leetcode

var CopyRandomList = copyRandomList

type tNode struct {
	Val    int
	Next   *tNode
	Random *tNode
}

func copyRandomList(head *tNode) *tNode {
	if head == nil {
		return head
	}

	m := make(map[*tNode]*tNode, 0)
	newHead := &tNode{}
	newQ := newHead

	oldQ := head
	for oldQ != nil {
		t := &tNode{
			Val: oldQ.Val,
		}
		m[oldQ] = t
		newQ.Next = t
		oldQ = oldQ.Next
		newQ = newQ.Next
	}

	oldQ = head
	newQ = newHead.Next
	for oldQ != nil {
		newQ.Random = m[oldQ.Random]
		oldQ = oldQ.Next
		newQ = newQ.Next
	}

	return newHead.Next
}
