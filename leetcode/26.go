package leetcode

func removeDuplicates1(nums []int) int {
	i, k := 0, 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != nums[i] { //new elements
			k++
			if k != j {
				nums[k] = nums[j]
			}
			i = j
		}
	}

	return k + 1
}

func removeDuplicates(nums []int) int {
	i, j := 0, 1
	for ; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
