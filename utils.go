package leetcode_go

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func sumInt(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
