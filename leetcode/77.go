package leetcode

var Combine = combine

func combine(n int, k int) (ans [][]int) {
	numbers := make([]int, 0)
	var helper func(level int)
	helper = func(level int) {
		if len(numbers)+(n-level+1) < k { //减枝
			return
		}
		if len(numbers) == k {
			comb := make([]int, k)
			copy(comb, numbers)
			ans = append(ans, comb)
			return
		}

		if level == n+1 {
			return
		}

		numbers = append(numbers, level)
		helper(level + 1)
		numbers = numbers[:len(numbers)-1]
		helper(level + 1)
	}

	helper(1)

	return
}
