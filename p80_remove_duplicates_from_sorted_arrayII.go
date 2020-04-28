package leetcode_go

func removeDuplicatesP80(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	pre, cur := 0, 1
	cnt := 1
	for cur < len(nums) {
		if nums[pre] == nums[cur] && cnt == 0 {
			cur++
		} else {
			if nums[pre] == nums[cur] {
				cnt--
			} else {
				cnt = 1
			}
			pre++
			nums[pre] = nums[cur]
			cur++
		}
	}
	return pre + 1
}
