package main

func removeElement(nums []int, val int) int {
	i := 0
	for i = 0; i < len(nums); i++ {
		if nums[i] == val {
			break
		}
	}

	for j := i + 1; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
