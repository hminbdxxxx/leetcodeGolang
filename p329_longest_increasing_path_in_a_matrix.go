package leetcode_go

var dirs [][]int = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	res := 1
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			res = max(res, helperP329(matrix, &dp, i, j))
		}
	}
	return res
}

func helperP329(matrix [][]int, dp *[][]int, i, j int) int {
	if (*dp)[i][j] != 0 {
		return (*dp)[i][j]
	}
	maxLen := 1
	for _, dir := range dirs {
		x, y := i+dir[0], j+dir[1]
		if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) || matrix[x][y] <= matrix[i][j] {
			continue
		}
		len := 1 + helperP329(matrix, dp, x, y)
		maxLen = max(maxLen, len)
	}
	(*dp)[i][j] = maxLen
	return (*dp)[i][j]
}
