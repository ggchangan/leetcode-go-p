package leetcode

var ReverseList = reverseList

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p, q := head, head.Next
	for q != nil {
		t := q.Next
		q.Next = p
		p = q
		q = t
	}

	return p
}
