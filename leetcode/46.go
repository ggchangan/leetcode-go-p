package leetcode

var Permute = permute

func permute(nums []int) (ans [][]int) {
	n := len(nums)
	var helper func(level int)
	helper = func(level int) {
		if level == n {
			perm := make([]int, n)
			copy(perm, nums)
			ans = append(ans, perm)
		}

		for i := level; i < n; i++ {
			nums[level], nums[i] = nums[i], nums[level]
			helper(level + 1)
			nums[level], nums[i] = nums[i], nums[level]
		}
	}
	helper(0)

	return
}
