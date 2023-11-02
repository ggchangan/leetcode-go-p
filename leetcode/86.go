package leetcode

var Partition = partition

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	smallHead := &ListNode{}
	small := smallHead
	bigHead := &ListNode{}
	big := bigHead
	p := head
	for p != nil {
		if p.Val < x {
			small.Next = p
			small = small.Next
		} else {
			big.Next = p
			big = big.Next
		}
		p = p.Next
	}
	small.Next = bigHead.Next
	big.Next = nil

	return smallHead.Next
}
