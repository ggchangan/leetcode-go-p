package leetcode

var TrailingZeroes = trailingZeroes

func trailingZeroes1(n int) (ans int) {
	for i := 0; i <= n; i += 5 {
		x := i
		for x != 0 && x%5 == 0 {
			ans++
			x /= 5
		}
	}

	return
}

func trailingZeroes(n int) (ans int) {
	for n > 0 {
		ans += n / 5
		n = n / 5
	}

	return
}
