package leetcode

var ProductExceptSelf = productExceptSelf

func productExceptSelf(nums []int) (ans []int) {
	n := len(nums)
	L := make([]int, n)
	R := make([]int, n)

	L[0] = 1
	R[n-1] = 1
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
		R[n-1-i] = R[n-i] * nums[n-i]
	}
	ans = make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = L[i] * R[i]
	}

	return
}
