package leetcode

var DeleteDuplicates = deleteDuplicates

func deleteDuplicates(head *ListNode) *ListNode {
	q, p := head, head
	var preQ *ListNode

	for p != nil {
		cnt := 0
		for p != nil && p.Val == q.Val {
			cnt++
			p = p.Next
		}

		//有重复
		if cnt > 1 {
			if preQ == nil {
				head = p
			} else {
				preQ.Next = p
			}
		} else {
			preQ = q
		}

		q = p
	}

	return head
}
