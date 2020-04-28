package leetcode_go

func numberOfArithmeticSlices(A []int) int {
	dp := make([]int, len(A))
	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] == A[i-1]-A[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	res := 0
	for _, n := range dp {
		res += n
	}
	return res
}
