package leetcode_go

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	backtrackP47([]int{}, nums, &res)
	return res
}

func backtrackP47(curPath []int, choices []int, res *[][]int) {
	visited := make(map[int]bool)
	for _, num := range choices {
		visited[num] = false
	}

	if len(choices) == 0 {
		*res = append(*res, curPath)
		return
	}

	for idx, num := range choices {
		if visited[num] {
			continue
		}
		visited[num] = true
		curPath = append(append([]int{}, curPath...), num)
		nextChoices := append([]int{}, choices[:idx]...)
		nextChoices = append(nextChoices, choices[idx+1:]...)
		backtrackP47(curPath, nextChoices, res)
		curPath = curPath[:len(curPath)-1]
	}
}
