package leetcode

func removeDuplicates80(nums []int) int {
	//k points to next add element position
	i, k, j := 0, 0, 0
	for ; j < len(nums); j++ {
		if nums[j] != nums[i] { //new elements
			nums[k] = nums[i]
			k++
			if j-i >= 2 {
				nums[k] = nums[i]
				k++
			}
			i = j
		}
	}

	nums[k] = nums[i]
	k++
	if j-i >= 2 {
		nums[k] = nums[i]
		k++
	}

	return k
}
