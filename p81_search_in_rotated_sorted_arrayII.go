package leetcode_go

func searchP81(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < nums[right] {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if nums[mid] < nums[right] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			right--
		}
	}
	return false
}
