package leetcode

var MyPow = myPow

func myPow2(x float64, n int) float64 {
	var helper func(x float64, n int) float64
	helper = func(x float64, n int) float64 {
		if n == 1 {
			return x
		}

		r := helper(x, n/2)
		if n%2 == 0 {
			return r * r
		} else {
			return x * r * r
		}
	}

	if n == 0 {
		return 1.0
	} else if n < 0 {
		return 1 / helper(x, -n)
	}

	return helper(x, n)

}

func myPow(x float64, n int) (ans float64) {
	ans = 1.0
	if n == 0 {
		return
	}

	sign := n < 0
	if sign {
		n = -n
	}

	t := x
	for n > 0 {
		l := n & 1
		if l == 1 {
			ans *= t
		}
		t *= t
		n >>= 1
	}

	if sign {
		ans = 1.0 / ans
	}

	return

}
