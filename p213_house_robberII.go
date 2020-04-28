package leetcode_go

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 0 {
		return 0
	}
	return max(rob2Helper(nums, 0, len(nums)-1), rob2Helper(nums, 1, len(nums)))
}

func rob2Helper(nums []int, l, r int) int {
	if r-l <= 1 {
		return nums[l]
	}

	dp := make([]int, r)
	dp[l] = nums[l]
	dp[l+1] = max(nums[l], nums[l+1])
	for i := l + 2; i < r; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(dp)-1]
}
