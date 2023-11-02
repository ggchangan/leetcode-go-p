package leetcode

// TODO more 练习
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				if dp[i-coins[j]]+1 < dp[i] {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
