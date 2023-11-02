package leetcode

var MySqrt = mySqrt

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	l, r := 1, x
	for l <= r {
		m := (l + r) >> 1
		v := m * m
		if x == v {
			return m
		} else if x > v {
			l = m + 1
		} else {
			r = m - 1
		}

	}

	return l - 1
}
