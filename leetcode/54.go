package leetcode

func spiralOrder(matrix [][]int) (ans []int) {
	m := len(matrix)
	n := len(matrix[0])
	ans = make([]int, m*n)
	u, d, l, r, c := 0, m-1, 0, n-1, 0
	for {
		for i := l; i <= r; i++ {
			ans[c] = matrix[u][i]
			c++
		}
		u++
		if u > d {
			break
		}
		for i := u; i <= d; i++ {
			ans[c] = matrix[i][r]
			c++
		}
		r--
		if r < l {
			break
		}
		for i := r; i >= l; i-- {
			ans[c] = matrix[d][i]
			c++
		}
		d--
		if d < u {
			break
		}
		for i := d; i >= u; i-- {
			ans[c] = matrix[i][l]
			c++
		}
		l++
		if l > r {
			break
		}
	}

	return
}
