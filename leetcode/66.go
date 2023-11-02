package leetcode

var PlusOne = plusOne

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		t := carry + digits[i]
		digits[i] = t % 10
		carry = t / 10
		if carry == 0 {
			break
		}
	}

	if carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
