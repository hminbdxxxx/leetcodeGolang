package leetcode_go

func getPermutation(n int, k int) string {
	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	res := ""
	f := make([]int, 10)
	f[0] = 1
	for i := 1; i <= 9; i++ {
		f[i] = f[i-1] * i
	}
	k--
	for i, tmp := n-1, k; i >= 0; i-- {
		j := tmp / f[i]
		res += nums[j]
		tmp %= f[i]
		nums = append(nums[:j], nums[j+1:]...)
	}
	return res
}
