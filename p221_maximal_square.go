package leetcode_go

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}

	res := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i == 0 || j == 0 {
				if matrix[i][j] == '1' {
					dp[i][j] = 1
				} else {
					dp[i][j] = 0
				}
			} else if matrix[i][j] == '1' {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			} else {
				dp[i][j] = 0
			}
			res = max(res, dp[i][j])
		}
	}
	return res
}
