package leetcode

var RotateRight = rotateRight

func rotateRight(head *ListNode, k int) *ListNode {
	//注意k==0的特殊情况
	if head == nil || k == 0 {
		return head
	}

	p := head
	n := 1
	for p.Next != nil {
		p = p.Next
		n++
	}
	tail := p
	k = k % n
	if k == 0 {
		return head
	}

	p = head
	var preP *ListNode

	for i := n - k; i > 0; i-- {
		preP = p
		p = p.Next
	}

	t := head
	head = p
	tail.Next = t
	preP.Next = nil

	return head
}
