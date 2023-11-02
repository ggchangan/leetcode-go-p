package leetcode

var HIndex = hIndex

func hIndex(citations []int) int {
	n := len(citations)
	counts := make([]int, n+1)
	for i := 0; i < n; i++ {
		number := citations[i]
		if number > n {
			number = n
		}
		counts[number] += 1
	}

	if counts[n] >= n {
		return n
	}

	for i := n - 1; i >= 0; i-- {
		counts[i] += counts[i+1]
		if counts[i] >= i {
			return i
		}
	}

	return 0
}
