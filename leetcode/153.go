package leetcode

var FindMin = findMin

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] <= nums[r] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return nums[l]
}
