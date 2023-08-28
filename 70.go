package main

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	s := 0
	f1 := 1
	f2 := 2

	for i := 2; i < n; i++ {
		s = f1 + f2
		f1 = f2
		f2 = s
	}

	return s

}
