package leetcode_go

type keyP494 struct {
	St, Sum int
}

func findTargetSumWays(nums []int, S int) int {
	m := make(map[keyP494]int)
	res := helperP494(nums, 0, S, &m)
	return res
}

func helperP494(nums []int, st, sum int, cached *map[keyP494]int) int {
	if st == len(nums) {
		if sum == 0 {
			return 1
		}
		return 0
	}
	if res, ok := (*cached)[keyP494{st, sum}]; ok {
		return res
	}
	c1 := helperP494(nums, st+1, sum-nums[st], cached)
	c2 := helperP494(nums, st+1, sum+nums[st], cached)
	(*cached)[keyP494{st, sum}] = c1 + c2
	return c1 + c2
}
