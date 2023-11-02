package leetcode

var ReverseBetween = reverseBetween

// 思路：分成3段拼接，第一段的fTail, 第二段的sHead, sTail, 第三段的tHead
// 易错点：当第一段为空的时候；fTail == nil, 此时head指向sHead会被第二段翻转，此时head应该被重新赋值为sTail
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	q := head
	var preQ *ListNode
	var fTail, sHead, sTail, tHead *ListNode
	for i := 1; i <= right; i++ {
		if left == i {
			fTail = preQ
			sHead = q
		}
		if right == i {
			sTail = q
			tHead = q.Next
		}
		preQ = q
		q = q.Next
	}

	//翻转中间链表
	if sTail != nil {
		sTail.Next = nil
	}
	reverse := func(head *ListNode) *ListNode {
		p, q := head, head.Next
		p.Next = nil
		for q != nil {
			t := q.Next
			q.Next = p
			p = q
			q = t
		}

		return p
	}
	reverse(sHead)

	if fTail != nil {
		fTail.Next = sTail
	} else {
		head = sTail
	}
	sHead.Next = tHead
	return head
}
