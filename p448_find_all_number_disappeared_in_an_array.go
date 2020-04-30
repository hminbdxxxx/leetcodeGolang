package leetcode_go

import (
	"math"
)

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		idx := int(math.Abs(float64(num)) - 1)
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}
	res := []int{}
	for i, num := range nums {
		if num > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
