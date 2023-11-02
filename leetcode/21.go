package leetcode

var MergeTwoLists = mergeTwoLists

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	p := list1
	q := list2

	var nHead *ListNode
	if p.Val < q.Val {
		nHead = p
		p = p.Next
	} else {
		nHead = q
		q = q.Next
	}

	n := nHead
	for p != nil && q != nil {
		if p.Val < q.Val {
			n.Next = p
			p = p.Next
		} else {
			n.Next = q
			q = q.Next
		}

		n = n.Next
	}

	if p != nil {
		n.Next = p
	}

	if q != nil {
		n.Next = q
	}

	return nHead
}
