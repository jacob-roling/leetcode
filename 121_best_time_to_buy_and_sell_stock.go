package leetcode

func BestTimeToBuyAndSellStock(prices []int) int {
	var max = 0
	for i, j := 0, 1; j < len(prices); {
		if prices[j]-prices[i] > max {
			max = prices[j] - prices[i]
		} else if prices[j] < prices[i] {
			i = j
		}
		j++
	}
	return max
}
