package leetcode

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	first := 1
	second := 0
	if s[1] != '0' {
		second += 1
	}
	if s[:2] <= "26" && s[:2] >= "10" {
		second += 1
	}

	for i := 2; i < len(s); i++ {
		a, b := first, second
		second = 0
		if s[i] != '0' {
			second = b
		}
		if s[i-1:i+1] <= "26" && s[i-1:i+1] >= "10" {
			second += a
		}

		first = b
	}

	return second
}
