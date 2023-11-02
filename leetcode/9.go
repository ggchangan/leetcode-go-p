package leetcode

var IsPalindromeNumber = isPalindromeNumber

func isPalindromeNumber(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	//当x小于或者等于revertedNumber时循环结束
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	return revertedNumber == x || revertedNumber/10 == x
}
