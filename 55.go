package main

func canJump(nums []int) bool {
	for i, m, n := 0, 0, len(nums); i <= m && i < n; i++ {
		if nums[i]+i > m {
			m = nums[i] + i
		}
		if m >= n-1 {
			return true
		}
	}
	return false
}
