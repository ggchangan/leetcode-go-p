package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

var HasCycle = hasCycle

func hasCycle(head *ListNode) bool {
	p, q := head, head
	for p != nil {
		p = p.Next
		if p == nil {
			return false
		}
		p = p.Next
		q = q.Next
		if p == q {
			return true
		}
	}

	return false
}
