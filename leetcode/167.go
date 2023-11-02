package leetcode

var TwoSum = twoSum

func twoSum(numbers []int, target int) []int {
	i, j := 0, 1
	for i < len(numbers) && j < len(numbers) {
		for j < len(numbers) && numbers[j]+numbers[i] < target {
			j++
		}

		if j < len(numbers) {
			if numbers[i]+numbers[j] == target {
				break
			}
		}
		i++
		j = i + 1
	}

	return []int{i + 1, j + 1}
}
