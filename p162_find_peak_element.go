package leetcode_go

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	var mid int
	for l < r {
		mid = (l + r) / 2
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return r
}
