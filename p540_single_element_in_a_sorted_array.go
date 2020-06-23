// https://leetcode.com/problems/single-element-in-a-sorted-array/

package leetcode_go

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l < r {
		mid = (l + r) / 2
		if nums[mid] == nums[mid^1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
