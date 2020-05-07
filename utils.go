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

func makeNIntSlice(n int, l int) []int {
	sli := make([]int, l)
	for i := 0; i < l; i++ {
		sli[i] = n
	}
	return sli
}

func sumInt(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
