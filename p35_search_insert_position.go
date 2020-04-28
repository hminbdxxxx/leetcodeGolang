package leetcode_go

func searchInsert(nums []int, target int) int {
	return bisearchP35(nums, target, 0, len(nums)-1)
}

func bisearchP35(nums []int, target int, l, r int) int {
	if l > r {
		return l
	}

	mid := (l + r) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return bisearchP35(nums, target, l, mid-1)
	} else {
		return bisearchP35(nums, target, mid+1, r)
	}
}
