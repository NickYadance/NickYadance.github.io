package codetop

func maxProfitV2(prices []int) int {
	dp := make([][]int, len(prices))
	dp[0] = []int{0, -prices[0]}
	max := 0
	for i := 1; i < len(prices); i++ {
		dp[i] = []int{Max(dp[i-1][0], dp[i-1][1]+prices[i]), Max(dp[i-1][1], dp[i-1][0]-prices[i])}
		if dp[i][0] > max {
			max = dp[i][0]
		}
	}
	return max
}
