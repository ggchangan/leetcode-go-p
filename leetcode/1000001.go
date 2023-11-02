package leetcode

//[0, 4], [1, 5], [3, 4], [9, 13]

func MergeData(data [][]int) (ans int) {
	maxHelper := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	l := len(data)
	if l <= 1 {
		ans = l
		return
	}
	stack := [][]int{data[0]}
	for i := 1; i < l; i++ {
		sL := len(stack)
		cur := stack[sL-1]
		if data[i][0] > cur[1] { //不合并
			stack = append(stack, data[i])
		} else { //合并
			stack[sL-1][1] = maxHelper(cur[1], data[i][1])
		}
	}
	//[0,5] [9, 13]
	//6

	//获取最大长度
	for i := 0; i < len(stack); i++ {
		cur := data[i][1] - data[i][0] + 1
		if ans < cur {
			ans = cur
		}
	}

	return
}
