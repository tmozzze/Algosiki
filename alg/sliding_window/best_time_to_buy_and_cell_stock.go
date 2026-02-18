package sliding_window

func MaxProfit(prices []int) (int, int) {
	maxProfit := 0

	day1 := 0
	day2 := 0

	left := 0
	right := 1

	for right < len(prices) {
		if maxProfit < (prices[right] - prices[left]) {
			maxProfit = prices[right] - prices[left]
			day1 = left
			day2 = right
		}

		if prices[right] < prices[left] {
			left = right
		}
		right++
	}

	return day1, day2
}

// 7,25,1,5,3,6,4
// l r
