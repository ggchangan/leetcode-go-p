package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	var l int
	var r int
	for ; k >= 0; k-- {
		if i < 0 {
			l = -1000000000 - 1
		} else {
			l = nums1[i]
		}
		if j < 0 {
			r = -1000000000 - 1
		} else {
			r = nums2[j]
		}

		if l > r {
			nums1[k] = l
			i--
		} else {
			nums1[k] = r
			j--
		}
	}
}
