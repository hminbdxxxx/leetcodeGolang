package leetcode_go

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pre, cur := 0, 1
	for cur < len(nums) {
		if nums[pre] == nums[cur] {
			cur++
		} else {
			pre++
			nums[pre] = nums[cur]
			cur++
		}
	}
	return pre + 1
}
