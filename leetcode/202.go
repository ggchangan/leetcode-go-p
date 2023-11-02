package leetcode

var IsHappy = isHappy

func isHappy(n int) bool {
	helper := func(n int) int {
		i := n
		sum := 0
		for i != 0 {
			t := i % 10
			sum += t * t
			i = i / 10
		}
		return sum
	}

	m := map[int]int{}
	for {
		new := helper(n)
		if _, ok := m[new]; ok {
			return false
		}
		if new == 1 {
			return true
		}

		m[n] = new
		n = new
	}
}
