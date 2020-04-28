package leetcode_go

func combine(n int, k int) [][]int {
	res := [][]int{}
	backtrackP77([]int{}, &res, 1, n, k)
	return res
}

func backtrackP77(curComs []int, res *[][]int, start, end, k int) {
	if k == 0 {
		*res = append(*res, curComs)
		return
	}

	for i := start; i <= end-k+1; i++ {
		curComs = append(append([]int{}, curComs...), i)
		backtrackP77(curComs, res, i+1, end, k-1)
		curComs = curComs[:len(curComs)-1]
	}
}
