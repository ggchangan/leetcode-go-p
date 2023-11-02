package leetcode

var MaxSubarraySumCircular = maxSubarraySumCircular

func maxSubarraySumCircular(nums []int) (ans int) {
	maxAns, minAns, maxDp, minDp, sum := nums[0], nums[0], nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if maxDp < 0 {
			maxDp = nums[i]
		} else {
			maxDp += nums[i]
		}
		if maxDp > maxAns {
			maxAns = maxDp
		}

		if minDp < 0 {
			minDp += nums[i]
		} else {
			minDp = nums[i]
		}

		if minDp < minAns {
			minAns = minDp
		}
		sum += nums[i]
	}
	if maxAns > sum-minAns || sum == minAns {
		ans = maxAns
	} else {
		ans = sum - minAns
	}

	return
}
