package leetcode

var MergeKLists = mergeKLists

// 自低向上分治
func mergeKLists(lists []*ListNode) *ListNode {
	merge := func(l1 *ListNode, l2 *ListNode) *ListNode {
		head := &ListNode{}
		p := head

		for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				p.Next = l1
				l1 = l1.Next
			} else {
				p.Next = l2
				l2 = l2.Next
			}
			p = p.Next
		}
		if l1 != nil {
			p.Next = l1
		}

		if l2 != nil {
			p.Next = l2
		}

		return head.Next
	}

	if len(lists) == 0 {
		return nil
	}

	interval := 1
	amount := len(lists)

	//二路归并排序，最后interval要小于amount
	for interval < amount {
		//i+interval < amount ==> i < amount - interval
		//间隔为interval时，两个组之间下标差interval *2
		for i := 0; i < amount-interval; i += interval * 2 {
			lists[i] = merge(lists[i], lists[i+interval])
		}
		interval *= 2
	}

	return lists[0]
}
