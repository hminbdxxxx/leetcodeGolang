package leetcode_go

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	helperP39(candidates, []int{}, target, 0, &res)
	return res
}

func helperP39(candidates []int, curSum []int, target int, start int, res *[][]int) {
	sum := sumInt(curSum)
	if sum > target {
		return
	} else if sum == target {
		(*res) = append(*res, curSum)
		return
	}

	for i := start; i < len(candidates); i++ {
		new := append(append([]int{}, curSum...), candidates[i])
		helperP39(candidates, new, target, i, res)
	}
}
