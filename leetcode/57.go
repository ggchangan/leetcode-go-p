package leetcode

var InsertInterval = insertInterval

func insertInterval(intervals [][]int, newInterval []int) (ans [][]int) {
	maxHelper := func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	l := len(intervals)
	if l == 0 {
		ans = append(ans, newInterval)
		return
	}

	ans = make([][]int, l+1)
	var i, cnt int
	for i = 0; i < l && intervals[i][0] < newInterval[0]; i++ {
		ans[cnt] = intervals[i]
		cnt++
	}
	if cnt == 0 || newInterval[0] > ans[cnt-1][1] { //cnt==0的特殊情况
		ans[cnt] = newInterval
		cnt++
	} else {
		ans[cnt-1][1] = maxHelper(ans[cnt-1][1], newInterval[1]) //注意下标，写程序的时候要画出来
	}

	for ; i < l; i++ {
		if intervals[i][0] > ans[cnt-1][1] {
			ans[cnt] = intervals[i]
			cnt++
		} else {
			ans[cnt-1][1] = maxHelper(ans[cnt-1][1], intervals[i][1])
		}
	}

	return ans[:cnt]
}
