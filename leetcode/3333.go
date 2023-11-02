package leetcode

func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}

	p, q, cur := l1, l2, head
	sum, carry := 0, 0
	for p != nil || q != nil {
		t := carry
		if p != nil {
			t += l1.Val
		}
		if q != nil {
			t += l2.Val
		}
		sum = t % 10
		carry = t / 10

		if p != nil {
			p.Val = sum
			cur.Next = p
			p = p.Next
		} else {
			q.Val = sum
			cur.Next = q
		}
		if q != nil {
			q = q.Next
		}
		cur = cur.Next
	}

	if carry == 1 {
		cur.Next = &ListNode{Val: 1}
	}

	return head.Next
}
