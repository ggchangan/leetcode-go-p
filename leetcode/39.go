package leetcode

var CombinationSum = combinationSum

func combinationSum(candidates []int, target int) (ans [][]int) {
	n := len(candidates)
	numbers := make([]int, 0)
	var helper func(level, target int)
	helper = func(level, target int) {
		if target == 0 {
			temp := make([]int, len(numbers))
			copy(temp, numbers)
			ans = append(ans, temp)
			return
		}

		if level == n {
			return
		}

		number := candidates[level]
		for i := 0; i <= target/number; i++ {
			for j := 0; j < i; j++ {
				numbers = append(numbers, number)
			}
			helper(level+1, target-number*i)
			numbers = numbers[:len(numbers)-i]
		}

	}
	helper(0, target)
	return
}
