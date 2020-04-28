package leetcode_go

func countBits(num int) []int {
	res := []int{0, 1}
	for i := 2; i <= num; i++ {
		var val int
		if i%2 == 0 {
			val = res[i/2]
		} else {
			val = res[i/2] + 1
		}
		res = append(res, val)
	}
	return res
}
