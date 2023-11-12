package main

func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	var profit int
	minPrice := prices[0]
	for _, price := range prices {
		if price-minPrice > profit {
			profit = price - minPrice
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return profit
}
