package leetcode

var SingleNumber = singleNumber

func singleNumber(nums []int) (ans int) {
	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}

	return
}
