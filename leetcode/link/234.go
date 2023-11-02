package link

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	data := []int{}
	slow, fast := head, head.Next
	data = append(data, slow.Val)
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		data = append(data, slow.Val)
	}
	if fast == nil { //奇数
		data = data[:len(data)-1]
	}

	p := slow.Next
	for i := len(data) - 1; i >= 0; i-- {
		if data[i] != p.Val {
			return false
		}
		p = p.Next
	}

	return true
}
