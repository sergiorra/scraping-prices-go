package price

func GetMinPrice(prices []int) int {
	var minPrice int
	for i, price := range prices {
		if i == 0 || price < minPrice {
			minPrice = price
		}
	}
	return minPrice
}
