package leetcode

var SingleNumber4 = singleNumber4

func singleNumber4(nums []int) (ans int) {
	l := len(nums)
	if l == 0 {
		return
	}
	// write code here
	ans = nums[0]
	for i := 1; i < len(nums); i++ {
		ans ^= nums[i]
	}

	return
}
