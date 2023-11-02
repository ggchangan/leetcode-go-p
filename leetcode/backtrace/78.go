package backtrace

var Subsets = subsets

func subsets(nums []int) (ans [][]int) {
	l := len(nums)
	if l == 0 {
		ans = append(ans, []int{})
		return
	}

	data := []int{}
	var dfs func(level int)
	dfs = func(level int) {
		if level == l {
			ans = append(ans, append([]int{}, data...))
			return
		}

		dfs(level + 1)
		data = append(data, nums[level])
		dfs(level + 1)
		data = data[:len(data)-1]
	}

	dfs(0)
	return
}
