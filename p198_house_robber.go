package leetcode_go

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}

	dp := []int{nums[0], max(nums[0], nums[1])}
	for i := 2; i < len(nums); i++ {
		dp = append(dp, max(dp[i-1], dp[i-2]+nums[i]))
	}
	return dp[len(dp)-1]
}
