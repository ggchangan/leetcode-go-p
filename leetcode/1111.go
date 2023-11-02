package leetcode

var MyAtoi1111 = myAtoi1111

func myAtoi1111(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	q := 0
	sign := 1
	for q < l && (s[q] < '1' || s[q] > '9') {
		if s[q] == '-' {
			sign = -1
		}
		q++
	}
	result := int32(0)
	for q < l && s[q] >= '0' && s[q] <= '9' {
		result = 10*result + int32(s[q]-'0')
		q++
	}

	result *= int32(sign)

	return int(result)
}
