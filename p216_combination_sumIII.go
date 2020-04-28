package leetcode_go

func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	curCom := []int{}
	backtrackP216(k, n, 1, curCom, &res)
	return res
}

func backtrackP216(k, n, start int, curCom []int, res *[][]int) {
	if k == 0 && n == 0 {
		*res = append(*res, curCom)
		return
	}
	if k == 0 || n == 0 {
		return
	}

	for i := start; i <= 9; i++ {
		curCom = append(append([]int{}, curCom...), i)
		backtrackP216(k-1, n-i, start+1, curCom, res)
		curCom = curCom[:len(curCom)-1]
	}
}
