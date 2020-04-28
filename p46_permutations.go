package leetcode_go

func permute(nums []int) [][]int {
	res := [][]int{}
	backtrackP46([]int{}, nums, &res)
	return res
}

func backtrackP46(curPath []int, choices []int, res *[][]int) {
	if len(choices) == 0 {
		*res = append(*res, curPath)
		return
	}

	for idx, num := range choices {
		curPath = append(append([]int{}, curPath...), num)
		nextChoices := append([]int{}, choices[:idx]...)
		nextChoices = append(nextChoices, choices[idx+1:]...)
		backtrackP46(curPath, nextChoices, res)
		curPath = curPath[:len(curPath)-1]
	}
}
