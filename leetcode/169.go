package leetcode

func majorityElement(nums []int) int {
	cnt := 1
	number := nums[0]
	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			number = nums[i]
			cnt = 1
		} else {
			if nums[i] == number {
				cnt++
			} else {
				cnt--
			}
		}
	}

	return number
}
