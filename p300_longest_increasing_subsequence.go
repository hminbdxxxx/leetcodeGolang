package leetcode_go

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		maxLen := 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxLen = max(maxLen, dp[j]+1)
			}
		}
		dp[i] = maxLen
	}

	res := 0
	for _, n := range dp {
		res = max(res, n)
	}
	return res
}
