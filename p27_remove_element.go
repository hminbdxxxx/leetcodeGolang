package leetcode_go

func removeElement(nums []int, val int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i] = nums[res]
			res++
		}
	}
	return res
}
