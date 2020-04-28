package leetcode_go

func maxSumDivThree(nums []int) int {
	sums := []int{0, 0, 0}
	for _, num := range nums {
		sums2 := make([]int, 3)
		for i := 0; i < 3; i++ {
			sums2[i] = sums[i]
		}
		for _, i := range sums2 {
			sums[(i+num)%3] = max(sums[(i+num)%3], i+num)
		}
	}
	return sums[0]
}
