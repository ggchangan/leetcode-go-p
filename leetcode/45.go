package leetcode

var Jump = jump

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	max := nums[0]
	step := 1
	if max >= len(nums)-1 {
		return step
	}

	left := 1
	right := max
	for {
		tMax := max
		for i := left; i <= right; i++ {
			if i+nums[i] > tMax {
				tMax = i + nums[i]
			}
		}
		left = right + 1
		if tMax > max {
			max = tMax
			right = max
		}
		step++
		if max >= len(nums)-1 {
			return step
		}
	}
}
