package leetcode

var RemoveNthFromEnd = removeNthFromEnd

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, q := head, head
	var preQ *ListNode
	for i := 0; i < n; i++ {
		p = p.Next
	}
	for p != nil {
		p = p.Next
		preQ = q
		q = q.Next
	}
	if preQ == nil {
		return head.Next
	} else {
		preQ.Next = q.Next
		return head
	}
}
