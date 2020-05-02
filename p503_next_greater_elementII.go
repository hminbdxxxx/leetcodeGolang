package leetcode_go

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
	stk := []int{}
	var num int
	for i := 0; i < len(nums)*2; i++ {
		num = nums[i%len(nums)]
		for len(stk) != 0 && num > nums[stk[len(stk)-1]] {
			res[stk[len(stk)-1]] = num
			stk = stk[:len(stk)-1]
		}
		if i < len(nums) {
			stk = append(stk, i)
		}
	}
	return res
}
