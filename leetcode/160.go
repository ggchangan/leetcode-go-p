package leetcode

var GetIntersectionNode = getIntersectionNode

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	m := map[*ListNode]bool{}
	p, q := headA, headB
	for p != nil {
		m[p] = true
		p = p.Next
	}

	for q != nil {
		if m[q] {
			break
		}
		q = q.Next
	}

	return q
}
