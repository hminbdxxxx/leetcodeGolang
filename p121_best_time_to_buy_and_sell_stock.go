package leetcode_go

import (
	"math"
)

func maxProfit(prices []int) int {
	buy := math.MaxInt32
	res := 0
	for _, pri := range prices {
		buy = min(buy, pri)
		res = max(res, pri-buy)
	}
	return res
}
