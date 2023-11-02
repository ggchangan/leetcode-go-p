package leetcode

func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		tmp := nums[i]
		nums[i] = nums[n-i-1]
		nums[n-i-1] = tmp
	}
}
func rotate(nums []int, k int) {
	reverse(nums)
	s := k % len(nums)
	reverse(nums[:s])
	reverse(nums[s:])
}
