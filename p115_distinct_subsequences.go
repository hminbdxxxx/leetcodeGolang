package leetcode_go

func numDistinct(s string, t string) int {
	dp := make([][]int, len(s))
	for j := 0; j < len(s); j++ {
		dp[0] = append(dp[0], 1)
	}
	for i := 1; i < len(t); i++ {
		dp[i] = []int{0}
	}
	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}
