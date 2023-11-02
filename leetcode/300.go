package leetcode

var LengthOfLIS = lengthOfLIS

// 思路1：dp[i]代表以a[i]结尾的最长递增子序列长度
// dp[i] = max{dp[j] } + 1 if nums[i] > nums[j]
// return max{dp[i]}
func lengthOfLIS1(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = 1
	for i := 1; i < l; i++ {
		max := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > max {
				max = dp[j]
			}
		}
		dp[i] = max + 1
	}

	max := 0
	for i := 0; i < l; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}

// 思路2：dp[i]代表长度为i+1的最长严格递增子序列中的最小值
// dp[i]严格递增的，因为如果后面比前面小，则可更新前面值
// 更新dp数组的值，对于nums[i],如果超过前面任意一个值，则添加，否则，找到第一个大于nums[i]的dp
// 更新dp中的值为nums[i];因为nums[i] > 前面dp的最小值，可将此值组成一个前面最大长度中最后一个
// 最终返回len(dp)
func lengthOfLIS(nums []int) int {
	l := len(nums)
	dp := []int{nums[0]}
	helper := func(l, r, target int) int {
		for l <= r {
			m := (l + r) >> 1
			if dp[m] == target {
				return m
			} else if dp[m] > target {
				r = m - 1
			} else {
				l = m + 1
			}
		}

		return l
	}
	for i := 1; i < l; i++ {
		j := helper(0, len(dp)-1, nums[i])
		if j == len(dp) {
			dp = append(dp, nums[i])
		} else {
			dp[j] = nums[i]
		}
	}

	return len(dp)
}
