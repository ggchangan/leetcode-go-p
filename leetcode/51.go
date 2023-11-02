package leetcode

import (
	"strings"
)

func solveNQueens(n int) (ans [][]string) {
	all := make([][]int, 0)
	col := make([]bool, n)
	pie := make([]bool, 2*n)
	na := make([]bool, 2*n)
	var helper func(int, []int)
	helper = func(l int, one []int) {
		if l == n {
			copied := make([]int, len(one)) //这地方特别容易出错，helper直接访问slice，内部对slice的改变最终会反应到原始slice上，原始slice是全局的，那么只有一个结果
			copy(copied, one)
			all = append(all, copied)
			return
		}

		for i := 0; i < n; i++ {
			naIndex := l - i + n //row-col and then shift，这地方刚开始na使用绝对值来索引，是错误的（1，3）（2，0）是可以放置的
			if col[i] || pie[l+i] || na[naIndex] {
				continue
			}

			col[i] = true
			pie[l+i] = true
			na[naIndex] = true
			one[l] = i

			helper(l+1, one)

			col[i] = false
			pie[l+i] = false
			na[naIndex] = false
			one[l] = 0

		}
	}

	output := func() (ans [][]string) {
		for i := 0; i < len(all); i++ {
			one := make([]string, n)
			for j := 0; j < len(all[i]); j++ {
				row := make([]string, n)
				for k := 0; k < n; k++ {
					row[k] = "."
				}
				row[all[i][j]] = "Q"
				one[j] = strings.Join(row, "")
			}
			ans = append(ans, one)
		}
		return
	}

	one := make([]int, n)
	helper(0, one)
	ans = output()

	return
}

//[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
