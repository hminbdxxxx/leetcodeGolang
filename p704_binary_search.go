// https://leetcode.com/problems/binary-search/

package leetcode_go

func searchP704(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
