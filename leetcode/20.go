package leetcode

var IsValid = isValid

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, 0)
	for i := range s {
		if m[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
