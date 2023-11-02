package leetcode

import "sort"

var SearchMatrix = searchMatrix

// 将整个矩阵看成1个递增数组
func searchMatrix2(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	l, r := 0, m*n-1
	for l <= r {
		mid := (r + l) / 2
		v := matrix[mid/n][mid%n]
		if v == target {
			return true
		} else if v < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}

// 两次二分搜索
func searchMatrix1(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	l, r := 0, m-1
	for l <= r {
		m := (r + l) / 2
		v := matrix[m][0]
		if v == target {
			return true
		} else if v < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	row := l - 1
	if row < 0 {
		return false
	}

	n := len(matrix[0])
	l, r = 0, n-1
	for l <= r {
		m := (r + l) / 2
		v := matrix[row][m]
		if v == target {
			return true
		} else if v < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return false
}

// 一次搜索简化版
func searchMatrix3(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i := sort.Search(m*n, func(i int) bool {
		return matrix[i/n][i%n] >= target
	})

	return i < m*n && matrix[i/n][i%n] == target
}

// 二次搜索简化版本
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i := sort.Search(m, func(i int) bool {
		return matrix[i][0] >= target
	})
	if i < m && matrix[i][0] == target {
		return true
	}

	if i == 0 {
		return false
	}

	row := i - 1
	i = sort.Search(n, func(i int) bool {
		return matrix[row][i] >= target
	})

	return i < n && matrix[row][i] == target
}
