package link

var SwapPairs = swapPairs

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p, head1, head2 := head, &ListNode{}, &ListNode{}
	p1, p2 := head1, head2

	for p != nil && p.Next != nil {
		p1.Next = p
		p1 = p1.Next

		p2.Next = p.Next
		p2 = p2.Next
		p = p.Next.Next
	}

	if p != nil {
		p1.Next = p
		p1 = p1.Next
	}
	p1.Next = nil
	p2.Next = nil

	p1, p2 = head1.Next, head2.Next
	head3 := &ListNode{}
	p = head3
	for p2 != nil && p1 != nil {
		p.Next = p2
		p = p.Next
		p2 = p2.Next
		p.Next = p1
		p = p.Next
		p1 = p1.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	return head3.Next
}
