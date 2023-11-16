package codetop

func maxProfit(prices []int) int {
	max := 0
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		max = Max(max, prices[i]-min)
	}
	return max
}
