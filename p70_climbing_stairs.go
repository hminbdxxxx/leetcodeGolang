package leetcode_go

func climbStairs(n int) int {
	memo := make([]int, n)
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}
