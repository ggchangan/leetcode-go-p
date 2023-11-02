package leetcode

var AddTwoNumbers = addTwoNumbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	var p *ListNode

	carry := 0
	for l1 != nil && l2 != nil {
		n1 := &ListNode{
			Val: (l1.Val + l2.Val + carry) % 10,
		}
		carry = (l1.Val + l2.Val + carry) / 10
		if l3 == nil {
			l3 = n1
			p = l3
		} else {
			p.Next = n1
			p = p.Next
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	head := l1
	if head == nil {
		head = l2
	}

	for head != nil {
		n1 := &ListNode{
			Val: (head.Val + carry) % 10,
		}
		carry = (head.Val + carry) / 10
		p.Next = n1

		p = p.Next
		head = head.Next
	}

	if carry != 0 {
		p.Next = &ListNode{
			Val: 1,
		}
	}

	return l3
}
