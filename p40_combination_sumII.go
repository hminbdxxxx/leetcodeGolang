package leetcode_go

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)
	helperP40(candidates, []int{}, target, 0, &res)
	return res
}

func helperP40(candidates []int, curSum []int, target int, start int, res *[][]int) {
	sum := sumInt(curSum)
	if sum > target {
		return
	} else if sum == target {
		(*res) = append(*res, curSum)
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		new := append(append([]int{}, curSum...), candidates[i])
		helperP40(candidates, new, target, i+1, res)
	}
}
