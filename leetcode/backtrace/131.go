package backtrace

var Partition = partition

func partition(str string) (ans [][]string) {
	sL := len(str)
	if sL == 1 {
		ans = append(ans, []string{str})
		return
	}

	dp := make([][]bool, sL)
	for i := 0; i < sL; i++ {
		dp[i] = make([]bool, sL)
		dp[i][i] = true
		if i+1 < sL && str[i+1] == str[i] {
			dp[i][i+1] = true
		}
	}
	for d := 2; d < sL; d++ { //对角线填充
		for s := 0; s < sL-d; s++ {
			e := s + d
			dp[s][e] = str[s] == str[e] && dp[s+1][e-1]
		}
	}

	data := make([]string, 0)
	var dfs func(level int)
	dfs = func(level int) {
		if level == sL {
			ans = append(ans, append([]string{}, data...))
			return
		}

		i := level
		for j := i; j < sL; j++ {
			if dp[i][j] {
				data = append(data, str[i:j+1])
				dfs(j + 1)
				data = data[:len(data)-1]
			}
		}

	}

	dfs(0)

	return
}
