package leetcode_go

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	helperP90(nums, []int{}, 0, &res)
	return res
}

func helperP90(nums []int, curSet []int, pos int, res *[][]int) {
	*res = append(*res, curSet)
	for i := pos; i < len(nums); i++ {
		curSet = append(append([]int{}, curSet...), nums[i])
		helperP90(nums, curSet, i+1, res)
		curSet = curSet[:len(curSet)-1]
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
}
