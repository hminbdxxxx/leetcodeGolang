package leetcode_go

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	var left, right, mid int
	left, right = 0, len(matrix)-1
	for left < right {
		mid = (left + right) / 2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	var row int
	if right > 0 {
		row = right - 1
	} else {
		row = right
	}
	left, right = 0, len(matrix[0])-1
	for left < right {
		mid = (left + right) / 2
		if matrix[row][mid] == target {
			return true
		}
		if matrix[row][mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
