package leetcode_go

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	row, column := len(grid)-1, len(grid[0])-1
	dp := make([][]int, row+1)
	for i := range dp {
		dp[i] = make([]int, column+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}

		}
	}

	return dp[row][column]
}
