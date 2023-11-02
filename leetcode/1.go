package leetcode

var TwoSum2 = twoSum2

func twoSum2(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[target-nums[i]]; ok {
			return []int{m[target-nums[i]], i}
		}
		m[nums[i]] = i
	}

	return []int{0, 0}
}
