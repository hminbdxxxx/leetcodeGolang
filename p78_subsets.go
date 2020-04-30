package leetcode_go

import (
	"sort"
)

func subsets(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{[]int{}}
	for _, num := range nums {
		toAdd := [][]int{}
		for _, set := range res {
			new := append(append([]int{}, set...), num)
			toAdd = append(toAdd, new)
		}
		res = append(res, toAdd...)
	}
	return res
}
