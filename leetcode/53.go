package leetcode

// dp[i] = max(dp[i-1],0) + nums[i] dp[i]以a[i]结尾的最大连续子数组之和
func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	//get all dp
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	m := dp[0]
	for i := 0; i < len(dp); i++ {
		if dp[i] > m {
			m = dp[i]
		}
	}

	return m
}
