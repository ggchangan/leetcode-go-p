package leetcode

var SortList = sortList

func sortList(head *ListNode) *ListNode {
	findMiddle := func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		slow, fast := head, head
		for fast != nil && fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		}

		return slow
	}

	merge := func(left *ListNode, right *ListNode) *ListNode {
		if left == nil && right == nil {
			return nil
		}

		if left == nil {
			return right
		}

		if right == nil {
			return left
		}

		var head *ListNode
		var p *ListNode
		if left.Val <= right.Val {
			head = left
			p = head
			left = left.Next
		} else {
			head = right
			p = head
			right = right.Next
		}

		for left != nil && right != nil {
			if left.Val <= right.Val {
				p.Next = left
				left = left.Next
			} else {
				p.Next = right
				right = right.Next
			}
			p = p.Next
		}

		if left != nil {
			p.Next = left
		}

		if right != nil {
			p.Next = right
		}

		return head
	}

	var sort func(head *ListNode) *ListNode
	sort = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}

		mid := findMiddle(head)
		t := mid.Next
		mid.Next = nil
		left := sort(head)
		right := sort(t)
		return merge(left, right)
	}

	return sort(head)
}
