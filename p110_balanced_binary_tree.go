package leetcode_go

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helperP110(root) != -1
}

func helperP110(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := helperP110(root.Left), helperP110(root.Right)
	if left == -1 {
		return -1
	}
	if right == -1 {
		return -1
	}
	if math.Abs(float64(left-right)) > 1 {
		return -1
	}
	return max(left, right) + 1
}
